package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

func QueryTicket(ctx context.Context, condition map[string]interface{}) ([]*model.TicketDO, error) {
	var tickets []*model.TicketDO
	query := configs.DB.Model(&model.TicketDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	query = query.Where("del_flag = ?", 0).Order("create_time desc")
	if err := query.Scan(&tickets).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("query ticket failed, err: %v", err)
		return nil, err
	}
	return tickets, nil
}

// BatchInsertTicket 批量插入车票
func BatchInsertTicket(ctx context.Context, tickets []*model.TicketDO) error {
	err := configs.DB.Create(&tickets).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("batch insert ticket failed, err: %v", err)
		return err
	}
	return nil
}

// CancelTicket 取消车票
func CancelTicket(ctx context.Context, ticketID int64) error {
	err := configs.DB.Model(&model.TicketDO{}).Where("id = ?", ticketID).Update("ticket_status", 1).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("cancel ticket failed, err: %v", err)
		return err
	}
	return nil
}
