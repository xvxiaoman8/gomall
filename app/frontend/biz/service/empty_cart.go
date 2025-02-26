package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xvxiaoman8/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/xvxiaoman8/gomall/app/frontend/utils"
	rpccart "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp *common.Empty, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{UserId: userId})

	return
}
