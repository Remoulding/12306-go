package service

import (
	"errors"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/req"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/resp"
	. "github.com/Remoulding/12306-go/user-service/biz/model"
	"gorm.io/gorm"
)

// UserService 用户信息服务接口
type UserService interface {
	// QueryUserByUserId 根据用户 ID 查询用户信息
	QueryUserByUserId(userId string) (*UserQueryRespDTO, error)

	// QueryUserByUsername 根据用户名查询用户信息
	QueryUserByUsername(username string) (*UserQueryRespDTO, error)

	// QueryUserDeletionNum 根据证件类型和证件号查询注销次数
	QueryUserDeletionNum(idType int, idCard string) (int, error)

	// Update 根据用户 ID 修改用户信息
	Update(requestParam *UserUpdateReqDTO) error
}

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

func (s *UserServiceImpl) QueryUserByUserId(userId string) (*UserQueryRespDTO, error) {
	var user UserDO
	if err := s.db.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在，请检查用户ID是否正确")
		}
		return nil, err
	}
	// Convert UserDO to UserQueryRespDTO
	return &UserQueryRespDTO{ /* populate fields */ }, nil
}

func (s *UserServiceImpl) QueryUserByUsername(username string) (*UserQueryRespDTO, error) {
	var user UserDO
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在，请检查用户名是否正确")
		}
		return nil, err
	}
	// Convert UserDO to UserQueryRespDTO
	return &UserQueryRespDTO{ /* populate fields */ }, nil
}

func (s *UserServiceImpl) QueryUserDeletionNum(idType int, idCard string) (int, error) {
	var count int64
	if err := s.db.Model(&UserDeletionDO{}).Where("id_type = ? AND id_card = ?", idType, idCard).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *UserServiceImpl) Update(requestParam *UserUpdateReqDTO) error {
	var user UserDO
	if err := s.db.Where("username = ?", requestParam.Username).First(&user).Error; err != nil {
		return err
	}
	// Update user information
	if err := s.db.Model(&user).Updates(UserDO{ /* populate fields from requestParam */ }).Error; err != nil {
		return err
	}
	// Update user mail if necessary
	if requestParam.Mail != "" && requestParam.Mail != user.Mail {
		if err := s.db.Where("mail = ?", user.Mail).Delete(&UserMailDO{}).Error; err != nil {
			return err
		}
		userMail := UserMailDO{
			Mail:     requestParam.Mail,
			Username: requestParam.Username,
		}
		if err := s.db.Create(&userMail).Error; err != nil {
			return err
		}
	}
	return nil
}
