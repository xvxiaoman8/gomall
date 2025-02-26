package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/xvxiaoman8/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/xvxiaoman8/gomall/app/frontend/utils"
	rpcorder "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
)

type DeleteOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteOrderService(Context context.Context, RequestContext *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteOrderService) Run(OrderId string, req *order.DeleteOrderReq) (resp *common.Empty, err error) {
	// 获取当前用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	rpcReq := &rpcorder.DeleteOrderReq{
		OrderId: OrderId,
		UserId:  userId,
	}
	//打印
	_, err = rpc.OrderClient.DeleteOrder(h.Context, rpcReq)
	return
}
