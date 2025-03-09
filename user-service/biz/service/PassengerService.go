package service

import (
	"context"
	"strconv"

	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/user-service/biz/dao"
	"github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/Remoulding/12306-go/user-service/configs"
	"github.com/Remoulding/12306-go/user-service/tools"
	"gorm.io/gorm"
)

// PassengerServiceImpl 乘车人服务实现
type PassengerServiceImpl struct {
	db *gorm.DB
}

// NewPassengerServiceImpl 创建乘车人服务实现
func NewPassengerServiceImpl() *PassengerServiceImpl {
	return &PassengerServiceImpl{db: configs.DB}
}

func (s *PassengerServiceImpl) ListPassengerQueryByUsername(ctx context.Context, req *user_service.ListPassengerByUsernameReq) (*user_service.ListPassengerByUsernameResp, error) {
	resp := &user_service.ListPassengerByUsernameResp{}
	passengers, err := dao.QueryPassengersByUsername(ctx, req.GetUsername())
	if err != nil {
		log.WithContext(ctx).Errorf("PassengerServiceImpl.ListPassengerQueryByUsername failed, err: %v", err)
		resp.Message = "查询失败"
		return resp, nil
	}
	data := make([]*user_service.Passenger, 0)
	for _, passenger := range passengers {
		data = append(data, &user_service.Passenger{
			// TODO
			Id:           strconv.Itoa(int(passenger.ID)),
			Username:     passenger.Username,
			RealName:     passenger.RealName,
			IdType:       int32(passenger.IDType),
			IdCard:       passenger.IDCard,
			DiscountType: int32(passenger.DiscountType),
			Phone:        passenger.Phone,
			VerifyStatus: int32(passenger.VerifyStatus),
		})
	}
	resp.Data = data
	resp.Success = true
	return resp, nil
}

func (s *PassengerServiceImpl) ListPassengerQueryByIds(ctx context.Context, req *user_service.ListPassengerByIdsReq) (*user_service.ListPassengerByIdsResp, error) {
	resp := &user_service.ListPassengerByIdsResp{}
	passengers, err := dao.QueryPassengersByUsername(ctx, req.GetUsername())
	if err != nil {
		log.WithContext(ctx).Errorf("PassengerServiceImpl.ListPassengerQueryByIds failed, err: %v", err)
		resp.Message = "查询失败"
		return resp, nil
	}
	mp := map[int64]struct{}{}
	for _, id := range req.GetIds() {
		mp[id] = struct{}{}
	}
	data := make([]*user_service.PassengerActualData, 0)
	for _, passenger := range passengers {
		if _, ok := mp[passenger.ID]; ok {
			data = append(data, &user_service.PassengerActualData{
				Id:           strconv.Itoa(int(passenger.ID)),
				Username:     passenger.Username,
				RealName:     passenger.RealName,
				IdType:       int32(passenger.IDType),
				IdCard:       passenger.IDCard,
				DiscountType: int32(passenger.DiscountType),
				Phone:        passenger.Phone,
				VerifyStatus: int32(passenger.VerifyStatus),
			})
		}
	}
	resp.Data = data
	resp.Success = true
	return resp, nil
}

func (s *PassengerServiceImpl) SavePassenger(ctx context.Context, req *user_service.PassengerReq) (*user_service.SavePassengerResp, error) {
	resp := &user_service.SavePassengerResp{}
	id, _ := strconv.ParseInt(req.GetId(), 10, 64)
	passengerDO := &model.PassengerDO{
		ID:           id,
		RealName:     req.GetRealName(),
		IDType:       int(req.GetIdType()),
		IDCard:       req.GetIdCard(),
		DiscountType: int(req.GetDiscountType()),
		Phone:        req.GetPhone(),
	}
	if err := dao.InsertPassenger(ctx, passengerDO); err != nil {
		log.WithContext(ctx).Errorf("PassengerServiceImpl.SavePassenger failed, err: %v", err)
		resp.Message = "保存失败"
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

func (s *PassengerServiceImpl) UpdatePassenger(ctx context.Context, req *user_service.PassengerReq) (*user_service.UpdatePassengerResp, error) {
	resp := &user_service.UpdatePassengerResp{}
	id, _ := strconv.ParseInt(req.GetId(), 10, 64)
	passengerDO := &model.PassengerDO{
		ID:           id,
		Username:     tools.ContextGetUserInfo(ctx, configs.UserNameKey),
		RealName:     req.GetRealName(),
		IDType:       int(req.GetIdType()),
		IDCard:       req.GetIdCard(),
		DiscountType: int(req.GetDiscountType()),
		Phone:        req.GetPhone(),
	}

	if err := dao.UpdatePassenger(ctx, passengerDO); err != nil {
		log.WithContext(ctx).Errorf("PassengerServiceImpl.UpdatePassenger failed, err: %v", err)
		resp.Message = "更新失败"
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

func (s *PassengerServiceImpl) RemovePassenger(ctx context.Context, req *user_service.RemovePassengerReq) (*user_service.RemovePassengerResp, error) {
	resp := &user_service.RemovePassengerResp{}
	if err := dao.DeletePassengers(ctx, req.GetId()); err != nil {
		log.WithContext(ctx).Errorf("PassengerServiceImpl.RemovePassenger failed, err: %v", err)
		resp.Message = "删除失败"
	}
	resp.Success = true
	return resp, nil
}
