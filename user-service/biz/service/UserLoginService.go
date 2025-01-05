package service

import (
	"context"
	"errors"
	"fmt"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/req"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/resp"
	"github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/Remoulding/12306-go/user-service/utils"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type UserLoginService interface {
	Login(requestParam *UserLoginReqDTO) (*UserLoginRespDTO, error)
	CheckLogin(accessToken string) (*UserLoginRespDTO, error)
	Logout(accessToken string) error
	HasUsername(username string) bool
	Register(requestParam *UserRegisterReqDTO) (*UserRegisterRespDTO, error)
	Deletion(requestParam *UserDeletionReqDTO) error
}

type UserLoginServiceImpl struct {
	ctx               context.Context
	db                *gorm.DB
	redisClient       *redis.Client
	jwtSecretKey      string
	userRegisterCache *redis.Client
	userService       UserService
}

func NewUserLoginServiceImpl(ctx context.Context, db *gorm.DB, redisClient *redis.Client, jwtSecretKey string, userRegisterCache *redis.Client, userService UserService) *UserLoginServiceImpl {
	return &UserLoginServiceImpl{
		ctx:               ctx,
		db:                db,
		redisClient:       redisClient,
		jwtSecretKey:      jwtSecretKey,
		userRegisterCache: userRegisterCache,
		userService:       userService,
	}
}

func (s *UserLoginServiceImpl) Login(requestParam *UserLoginReqDTO) (*UserLoginRespDTO, error) {
	usernameOrMailOrPhone := requestParam.UsernameOrMailOrPhone
	var username string

	// 判断是否为邮箱
	if utils.IsEmail(usernameOrMailOrPhone) {
		userMail := &model.UserMailDO{}
		result := s.db.Where("mail = ?", usernameOrMailOrPhone).First(userMail)
		if result.Error != nil {
			return nil, errors.New("用户名/手机号/邮箱不存在")
		}
		username = userMail.Username
	} else {
		userPhone := &model.UserPhoneDO{}
		result := s.db.Where("phone = ?", usernameOrMailOrPhone).First(userPhone)
		if result.Error != nil {
			return nil, errors.New("手机号不存在")
		}
		username = userPhone.Username
	}

	// 查询用户信息
	user := &model.UserDO{}
	result := s.db.Where("username = ? AND password = ?", username, requestParam.Password).First(user)
	if result.Error != nil {
		return nil, errors.New("账号不存在或密码错误")
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"realName": user.RealName,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.jwtSecretKey))
	if err != nil {
		return nil, err
	}

	resp := &UserLoginRespDTO{
		UserId:      fmt.Sprintf("%d", user.ID),
		Username:    user.Username,
		RealName:    user.RealName,
		AccessToken: tokenString,
	}

	// 将登录信息存储到 Redis
	err = s.redisClient.Set(s.ctx, tokenString, resp, time.Minute*30).Err()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserLoginServiceImpl) CheckLogin(accessToken string) (*UserLoginRespDTO, error) {
	resp := &UserLoginRespDTO{}
	err := s.redisClient.Get(s.ctx, accessToken).Scan(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *UserLoginServiceImpl) Logout(accessToken string) error {
	err := s.redisClient.Del(s.ctx, accessToken).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *UserLoginServiceImpl) HasUsername(username string) bool {
	// TODO
	// 检查用户名是否存在于布隆过滤器中
	//if !s.userRegisterCache.Contains(username) {
	//	return false
	//}

	if exists, err := s.userRegisterCache.Do(s.ctx, "BF.EXISTS", "user_register", username).Bool(); err != nil && !exists {
		return false
	}

	// 检查 Redis 是否存在此用户名
	shardKey := "user_register_reuse_sharding" + strconv.Itoa(utils.HashShardingIdx(username))
	isMember, err := s.redisClient.SIsMember(s.ctx, shardKey, username).Result()
	if err != nil {
		logrus.Error("Error checking username in Redis", err)
		return false
	}

	return isMember
}

func (s *UserLoginServiceImpl) Register(requestParam *UserRegisterReqDTO) (*UserRegisterRespDTO, error) {
	// 用户名是否已注册
	if s.HasUsername(requestParam.Username) {
		return nil, errors.New("用户名已存在")
	}

	// 注册业务逻辑
	user := &model.UserDO{
		Username: requestParam.Username,
		Password: requestParam.Password,
		RealName: requestParam.RealName,
	}

	// 开始数据库事务
	tx := s.db.Begin()

	// 插入用户
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 插入手机号、邮箱等信息
	if requestParam.Phone != "" {
		userPhone := &model.UserPhoneDO{
			Phone:    requestParam.Phone,
			Username: requestParam.Username,
		}
		if err := tx.Create(userPhone).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if requestParam.Mail != "" {
		userMail := &model.UserMailDO{
			Mail:     requestParam.Mail,
			Username: requestParam.Username,
		}
		if err := tx.Create(userMail).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 提交事务
	tx.Commit()

	// 添加到布隆过滤器
	err := s.userRegisterCache.Do(s.ctx, "BF.ADD", "user_register", requestParam.Username).Err()
	if err != nil {
		log.Default().Println("Error adding username to Bloom Filter", err)
	}

	// 从 Redis 中删除
	shardKey := "user_register_reuse_sharding" + strconv.Itoa(utils.HashShardingIdx(requestParam.Username))
	s.redisClient.SRem(s.ctx, shardKey, requestParam.Username)

	return &UserRegisterRespDTO{
		Username: requestParam.Username,
		RealName: requestParam.RealName,
	}, nil
}

func (s *UserLoginServiceImpl) Deletion(requestParam *UserDeletionReqDTO) error {
	// 获取当前登录用户
	username := requestParam.Username
	user := &model.UserDO{}
	result := s.db.Where("username = ?", username).First(user)
	if result.Error != nil {
		return errors.New("用户不存在")
	}

	// 删除操作：删除用户信息、手机号、邮箱等
	tx := s.db.Begin()
	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除关联的手机号、邮箱等信息
	userPhone := &model.UserPhoneDO{}
	if err := tx.Where("username = ?", username).Delete(userPhone).Error; err != nil {
		tx.Rollback()
		return err
	}

	userMail := &model.UserMailDO{}
	if err := tx.Where("username = ?", username).Delete(userMail).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()

	// 删除 Redis 缓存
	err := s.redisClient.Del(s.ctx, username).Err()
	if err != nil {
		return err
	}

	return nil
}
