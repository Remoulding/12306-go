package service

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/idl-gen/user/user_service"
	. "github.com/Remoulding/12306-go/user-service/biz/model"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	user_service.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) user_service.UserServiceServer {
	return &UserServiceImpl{db: db}
}

func (s *UserServiceImpl) QueryUserByUserId(ctx context.Context, req *user_service.QueryUserByUserIdReq) (*user_service.QueryUserResp, error) {
	userId := req.GetUserId()
	var user UserDO
	if err := s.db.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在，请检查用户ID是否正确")
		}
		return nil, err
	}
	// Convert UserDO to UserQueryRespDTO
	return &user_service.QueryUserResp{}, nil
}

func (s *UserServiceImpl) QueryUserByUsername(ctx context.Context, req *user_service.QueryUserByUsernameReq) (*user_service.QueryUserResp, error) {
	username := req.GetUsername()
	var user UserDO
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在，请检查用户名是否正确")
		}
		return nil, err
	}
	// Convert UserDO to UserQueryRespDTO\\

	return &user_service.QueryUserResp{ /* populate fields */ }, nil
}

func (s *UserServiceImpl) QueryUserDeletionNum(ctx context.Context, req *user_service.QueryUserDeletionNumReq) (*user_service.QueryUserDeletionNumResp, error) {
	var count int64
	idType := req.GetIdType()
	idCard := req.GetIdCard()
	if err := s.db.Model(&UserDeletionDO{}).Where("id_type = ? AND id_card = ?", idType, idCard).Count(&count).Error; err != nil {
		return nil, err
	}
	return &user_service.QueryUserDeletionNumResp{DeletionNum: int32(count)}, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.UserUpdateResp, error) {
	// var user UserDO
	// username := req.GetUsername()
	// if err := s.db.Where("username = ?", requestParam.Username).First(&user).Error; err != nil {
	// 	return err
	// }
	// // Update user information
	// if err := s.db.Model(&user).Updates(UserDO{ /* populate fields from requestParam */ }).Error; err != nil {
	// 	return err
	// }
	// // Update user mail if necessary
	// if requestParam.Mail != "" && requestParam.Mail != user.Mail {
	// 	if err := s.db.Where("mail = ?", user.Mail).Delete(&UserMailDO{}).Error; err != nil {
	// 		return err
	// 	}
	// 	userMail := UserMailDO{
	// 		Mail:     requestParam.Mail,
	// 		Username: requestParam.Username,
	// 	}
	// 	if err := s.db.Create(&userMail).Error; err != nil {
	// 		return err
	// 	}
	// }
	return nil, nil
}
