package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Remoulding/12306-go/user-service/configs"
	json "github.com/bytedance/sonic"
	"time"

	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/user-service/biz/dao"
	"github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/Remoulding/12306-go/user-service/tools"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var log = configs.Log

type UserLoginServiceImpl struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUserLoginServiceImpl() *UserLoginServiceImpl {
	return &UserLoginServiceImpl{
		db:          configs.DB,
		redisClient: configs.UserLoginCache,
	}
}

func (s *UserLoginServiceImpl) Login(ctx context.Context, req *user_service.UserLoginReq) (*user_service.UserLoginResp, error) {
	resp := &user_service.UserLoginResp{}
	usernameOrMailOrPhone := req.UsernameOrMailOrPhone
	condition := map[string]interface{}{}
	// 判断是否为邮箱
	if tools.IsEmail(usernameOrMailOrPhone) {
		condition["mail"] = usernameOrMailOrPhone
	} else if tools.IsPhone(usernameOrMailOrPhone) {
		condition["phone"] = usernameOrMailOrPhone
	} else {
		condition["username"] = usernameOrMailOrPhone
	}

	// 查询用户信息
	user, err := dao.QueryUser(ctx, condition)
	if err != nil {
		log.WithContext(ctx).Errorf("UserLoginServiceImpl.Login failed, err: %v", err)
		resp.Message = "用户名或密码错误"
		return resp, nil
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"realName": user.RealName,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(configs.JwtSecretKey))
	data := &user_service.UserLoginData{
		UserId:      fmt.Sprintf("%d", user.ID),
		Username:    user.Username,
		RealName:    user.RealName,
		AccessToken: tokenString,
	}
	// 将登录信息存储到 Redis
	redisData, _ := json.MarshalString(data)
	err = s.redisClient.Set(ctx, tokenString, redisData, configs.CacheTTL).Err()
	if err != nil {
		resp.Message = "登录失败"
		return resp, nil
	}
	resp.Success = true
	resp.Data = data
	return resp, nil
}

func (s *UserLoginServiceImpl) CheckLogin(ctx context.Context, req *user_service.CheckLoginReq) (*user_service.UserLoginResp, error) {
	resp := &user_service.UserLoginResp{}
	// 获取 Redis 中的值
	redisData, err := s.redisClient.Get(ctx, req.AccessToken).Result()
	if err != nil {
		resp.Message = "校验失败"
		return resp, nil
	}
	// 解析 JSON 数据到结构体
	data := &user_service.UserLoginData{}
	err = json.Unmarshal([]byte(redisData), data)
	if err != nil {
		log.WithContext(ctx).Errorf("UserLoginServiceImpl.CheckLogin failed, err: %v", err)
		resp.Message = "校验失败"
		return resp, nil
	}
	resp.Success = true
	resp.Data = data
	return resp, nil
}
func (s *UserLoginServiceImpl) Logout(ctx context.Context, req *user_service.LogoutReq) (*user_service.LogoutResp, error) {
	resp := &user_service.LogoutResp{}
	err := s.redisClient.Del(ctx, req.AccessToken).Err()
	if err != nil {
		log.WithContext(ctx).Errorf("UserLoginServiceImpl.Logout failed, err: %v", err)
		resp.Message = "服务异常"
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

func (s *UserLoginServiceImpl) HasUsername(ctx context.Context, req *user_service.HasUsernameReq) (*user_service.HasUsernameResp, error) {
	resp := &user_service.HasUsernameResp{}
	// 查询DB用户名是否存在
	conditions := map[string]interface{}{"username": req.Username, "del_flag": 0}
	user, _ := dao.QueryUser(ctx, conditions)
	resp.Success = true
	resp.Exists = user != nil
	return resp, nil
}

func (s *UserLoginServiceImpl) Register(ctx context.Context, req *user_service.UserRegisterReq) (*user_service.UserRegisterResp, error) {
	resp := &user_service.UserRegisterResp{}
	user := &model.UserDO{
		Username:     req.GetUsername(),
		Password:     req.GetPassword(),
		RealName:     req.GetRealName(),
		Region:       req.GetRegion(),
		IDType:       int(req.GetIdType()),
		IDCard:       req.GetIdCard(),
		Phone:        req.GetPhone(),
		Telephone:    req.GetTelephone(),
		Mail:         req.GetMail(),
		UserType:     int(req.GetUserType()),
		VerifyStatus: int(req.GetVerifyState()),
		PostCode:     req.GetPostCode(),
		Address:      req.GetAddress(),
	}

	// 调用dao层注册
	err := dao.InsertUser(ctx, user)
	if err != nil {
		log.WithContext(ctx).Errorf("UserLoginServiceImpl.Register failed, err: %v", err)
		resp.Message = "注册失败"
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			resp.Message = "用户名已存在"
		}
		return resp, nil
	}
	resp.Success = true
	resp.Data = &user_service.UserRegisterData{
		Username: user.Username,
		RealName: user.RealName,
		Phone:    user.Phone,
	}

	return resp, nil
}

func (s *UserLoginServiceImpl) Deletion(ctx context.Context, req *user_service.UserDeletionReq) (*user_service.DeletionResp, error) {
	resp := &user_service.DeletionResp{}
	err := dao.DeleteUser(ctx, req.Username)
	if err != nil {
		log.WithContext(ctx).Errorf("UserLoginServiceImpl.Deletion failed, err: %v", err)
		resp.Message = "删除失败"
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}
