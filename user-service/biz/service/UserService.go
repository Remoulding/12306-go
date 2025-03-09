package service

import (
	"context"
	"errors"

	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/user-service/biz/dao"
	"github.com/Remoulding/12306-go/user-service/tools"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) QueryUserByUsername(ctx context.Context, req *user_service.QueryUserByUsernameReq) (*user_service.QueryUserResp, error) {
	resp := &user_service.QueryUserResp{}
	username := req.GetUsername()
	conditions := map[string]interface{}{"username": username, "del_flag": 0}
	user, err := dao.QueryUser(ctx, conditions)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Error("UserService.QueryUserByUsername failed, user not found")
			resp.Message = "用户不存在"
			return resp, nil
		}
		log.WithContext(ctx).Errorf("UserService.QueryUserByUsername failed, err: %v", err)
		resp.Message = "查询失败"
		return resp, nil
	}
	resp.Success = true
	resp.Data = tools.ConvertModelToQueryUserResp(user)
	return resp, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.UserUpdateResp, error) {
	resp := &user_service.UserUpdateResp{}
	userDo := tools.ConvertUserUpdateReqToModel(req)
	if err := dao.Update(ctx, userDo); err != nil {
		log.WithContext(ctx).Errorf("UserService.Update failed, err: %v", err)
		resp.Message = "更新失败"
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}
