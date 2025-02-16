package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/xvxiaoman8/gomall/app/order/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/order/biz/model"
	order "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	var cur_order model.Order
	if err := mysql.DB.Where("order_id = ? AND user_id = ?", req.OrderId, req.UserId).First(&cur_order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("order not found or user mismatch")
		}
		return nil, fmt.Errorf("failed to query order: %v", err)
	}

	// 2. 删除订单
	model.DeleteOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	resp = &order.DeleteOrderResp{
		Order: &order.OrderResult{
			OrderId: req.OrderId,
		},
	}
	return resp, nil
}
