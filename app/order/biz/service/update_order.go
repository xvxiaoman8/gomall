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

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// 1. 校验订单是否存在
	var cur_order model.Order
	if err := mysql.DB.Where("order_id = ? AND user_id = ?", req.OrderId, req.UserId).First(&cur_order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("order not found or user mismatch")
		}
		return nil, fmt.Errorf("failed to query order: %v", err)
	}

	// 2. 更新订单信息
	updates := make(map[string]interface{})
	if req.Address != nil {
		updates["country"] = req.Address.Country
		updates["state"] = req.Address.State
		updates["city"] = req.Address.City
		updates["street_address"] = req.Address.StreetAddress
		updates["zip_code"] = req.Address.ZipCode
		updates["updated_at"] = mysql.DB.NowFunc()
	}
	// 3. 执行更新
	if len(updates) > 0 {
		model.UpdateOrder(mysql.DB, s.ctx, req.UserId, req.OrderId, updates)
	}

	// 4. 返回响应
	resp = &order.UpdateOrderResp{
		Order: &order.OrderResult{
			OrderId: req.OrderId,
		},
	}
	return resp, nil
}
