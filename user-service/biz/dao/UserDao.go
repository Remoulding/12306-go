package dao

import (
	"context"
	"errors"

	"github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/Remoulding/12306-go/user-service/configs"
	"gorm.io/gorm"
)

var (
	db  = configs.DB
	log = configs.Log
)

func QueryUser(ctx context.Context, conditions map[string]interface{}) (*model.UserDO, error) {
	var user *model.UserDO
	query := db.Model(&model.UserDO{})
	for column, value := range conditions {
		query = query.Where(column+" = ?", value)
	}
	if err := query.First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Errorf("UserDao.QueryUser failed, err: %v", err)
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func Update(ctx context.Context, user *model.UserDO) error {
	if err := db.Model(&model.UserDO{}).Where("username = ?", user.Username).Updates(user).Error; err != nil {
		log.WithContext(ctx).Errorf("UserDao.Update failed, err: %v", err)
		return err
	}
	return nil
}

func InsertUser(ctx context.Context, user *model.UserDO) error {
	if err := db.Model(&model.UserDO{}).Create(user).Error; err != nil {
		log.WithContext(ctx).Errorf("UserDao.InsertUser failed, err: %v", err)
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, username string) error {
	// 惰性删除
	if err := db.Model(&model.UserDO{}).Where("username = ?", username).Update("del_flag", 1).Error; err != nil {
		log.WithContext(ctx).Errorf("UserDao.DeleteUser failed, err: %v", err)
		return err
	}
	return nil
}
