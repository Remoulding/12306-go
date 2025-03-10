package dao

import (
	"context"
	"github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/Remoulding/12306-go/user-service/configs"
)

func QueryPassengersByUsername(ctx context.Context, username string) ([]*model.PassengerDO, error) {
	passengers := make([]*model.PassengerDO, 0)
	if err := configs.DB.Model(&model.PassengerDO{}).Where("username = ?", username).Find(&passengers).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("PassengerDao.QueryPassengersByUsername failed, err: %v", err)
		return nil, err
	}
	return passengers, nil
}

func InsertPassenger(ctx context.Context, passenger *model.PassengerDO) error {
	if err := configs.DB.Model(&model.PassengerDO{}).Create(passenger).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("PassengerDao.InsertPassenger failed, err: %v", err)
		return err
	}
	return nil
}

func UpdatePassenger(ctx context.Context, passenger *model.PassengerDO) error {
	// id 和 name 相同
	if err := configs.DB.Model(&model.PassengerDO{}).Where("id = ? and username = ?", passenger.ID, passenger.Username).Updates(passenger).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("PassengerDao.UpdatePassenger failed, err: %v", err)
		return err
	}
	return nil
}

func DeletePassengers(ctx context.Context, ids ...int64) error {
	if err := configs.DB.Model(&model.PassengerDO{}).Where("id in (?)", ids).Update("del_flag", 1).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("PassengerDao.DeletePassengers failed, err: %v", err)
		return err
	}
	return nil
}
