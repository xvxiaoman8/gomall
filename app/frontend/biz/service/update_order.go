package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	order "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/xvxiaoman8/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/xvxiaoman8/gomall/app/frontend/utils"
	rpcorder "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderService) Run(OrderId string, req *order.UpdateOrderReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	rpcReq := &rpcorder.UpdateOrderReq{
		OrderId: OrderId,
		UserId:  userId,
		Address: &rpcorder.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	}
	// 调用RPC服务更新订单
	_, err = rpc.OrderClient.UpdateOrder(h.Context, rpcReq)
	if err != nil {
		return nil, fmt.Errorf("更新失败: %v", err)
	}

	return utils.H{
		"title": "Order",
	}, nil
}
