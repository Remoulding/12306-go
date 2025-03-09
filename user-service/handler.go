package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/user-service/biz/service"
)

type UserServiceHandler struct {
	user_service.UnimplementedUserServiceServer
}

func NewUserServiceHandler() user_service.UserServiceServer {
	return &UserServiceHandler{}
}

func (u UserServiceHandler) QueryUserByUsername(ctx context.Context, req *user_service.QueryUserByUsernameReq) (*user_service.QueryUserResp, error) {
	return service.NewUserServiceImpl().QueryUserByUsername(ctx, req)
}

func (u UserServiceHandler) Update(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.UserUpdateResp, error) {
	return service.NewUserServiceImpl().Update(ctx, req)

}

func (u UserServiceHandler) Login(ctx context.Context, req *user_service.UserLoginReq) (*user_service.UserLoginResp, error) {
	return service.NewUserLoginServiceImpl().Login(ctx, req)

}

func (u UserServiceHandler) CheckLogin(ctx context.Context, req *user_service.CheckLoginReq) (*user_service.UserLoginResp, error) {
	return service.NewUserLoginServiceImpl().CheckLogin(ctx, req)
}

func (u UserServiceHandler) Logout(ctx context.Context, req *user_service.LogoutReq) (*user_service.LogoutResp, error) {
	return service.NewUserLoginServiceImpl().Logout(ctx, req)

}

func (u UserServiceHandler) HasUsername(ctx context.Context, req *user_service.HasUsernameReq) (*user_service.HasUsernameResp, error) {
	return service.NewUserLoginServiceImpl().HasUsername(ctx, req)

}

func (u UserServiceHandler) Register(ctx context.Context, req *user_service.UserRegisterReq) (*user_service.UserRegisterResp, error) {
	return service.NewUserLoginServiceImpl().Register(ctx, req)

}

func (u UserServiceHandler) Deletion(ctx context.Context, req *user_service.UserDeletionReq) (*user_service.DeletionResp, error) {
	return service.NewUserLoginServiceImpl().Deletion(ctx, req)

}

func (u UserServiceHandler) ListPassengerQueryByUsername(ctx context.Context, req *user_service.ListPassengerByUsernameReq) (*user_service.ListPassengerByUsernameResp, error) {
	return service.NewPassengerServiceImpl().ListPassengerQueryByUsername(ctx, req)

}

func (u UserServiceHandler) ListPassengerQueryByIds(ctx context.Context, req *user_service.ListPassengerByIdsReq) (*user_service.ListPassengerByIdsResp, error) {
	return service.NewPassengerServiceImpl().ListPassengerQueryByIds(ctx, req)

}

func (u UserServiceHandler) SavePassenger(ctx context.Context, req *user_service.PassengerReq) (*user_service.SavePassengerResp, error) {
	return service.NewPassengerServiceImpl().SavePassenger(ctx, req)

}

func (u UserServiceHandler) UpdatePassenger(ctx context.Context, req *user_service.PassengerReq) (*user_service.UpdatePassengerResp, error) {
	return service.NewPassengerServiceImpl().UpdatePassenger(ctx, req)

}

func (u UserServiceHandler) RemovePassenger(ctx context.Context, req *user_service.RemovePassengerReq) (*user_service.RemovePassengerResp, error) {
	return service.NewPassengerServiceImpl().RemovePassenger(ctx, req)

}
