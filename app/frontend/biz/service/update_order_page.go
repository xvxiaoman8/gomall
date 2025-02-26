package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	order "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/xvxiaoman8/gomall/app/frontend/infra/rpc"
	"github.com/xvxiaoman8/gomall/app/frontend/types"
	frontendutils "github.com/xvxiaoman8/gomall/app/frontend/utils"
	rpcorder "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
)

type UpdateOrderPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderPageService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderPageService {
	return &UpdateOrderPageService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderPageService) Run(req *order.UpdateOrderPageReq) (resp map[string]any, err error) {
	// 获取当前用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	// 调用现有的ListOrder接口
	listResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: userId})
	if err != nil || listResp == nil {
		return nil, fmt.Errorf("获取订单列表失败")
	}

	// 从列表中找到目标订单
	var targetOrder *rpcorder.Order
	for _, o := range listResp.Orders {
		if o.OrderId == req.OrderId {
			targetOrder = o
			break
		}
	}
	if targetOrder == nil {
		return nil, fmt.Errorf("订单不存在")
	}

	return utils.H{
		"title": "修改订单",
		"order": &types.Order{
			OrderId: targetOrder.OrderId,
			Consignee: types.Consignee{
				Email:         targetOrder.Email,
				StreetAddress: targetOrder.Address.StreetAddress,
				City:          targetOrder.Address.City,
				State:         targetOrder.Address.State,
				Country:       targetOrder.Address.Country,
				ZipCode:       targetOrder.Address.ZipCode,
			},
		},
	}, nil
}
