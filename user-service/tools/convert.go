package tools

import (
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	. "github.com/Remoulding/12306-go/user-service/biz/model"
	"strconv"
)

func ConvertUserUpdateReqToModel(req *user_service.UserUpdateReq) *UserDO {
	id, _ := strconv.ParseInt(req.GetId(), 10, 64)
	return &UserDO{
		ID:       uint64(id),
		Username: req.GetUsername(),
		Mail:     req.GetMail(),
		UserType: int(req.GetUserType()),
		PostCode: req.GetPostCode(),
		Address:  req.GetAddress(),
	}
}

func ConvertModelToQueryUserResp(userDO *UserDO) *user_service.QueryUserData {
	return &user_service.QueryUserData{
		Username:     userDO.Username,
		RealName:     userDO.RealName,
		Region:       userDO.Region,
		IdType:       int32(userDO.IDType),
		IdCard:       userDO.IDCard,
		Phone:        userDO.Phone,
		Telephone:    userDO.Telephone,
		Mail:         userDO.Mail,
		UserType:     int32(userDO.UserType),
		VerifyStatus: int32(userDO.VerifyStatus),
		PostCode:     userDO.PostCode,
		Address:      userDO.Address,
	}
}
