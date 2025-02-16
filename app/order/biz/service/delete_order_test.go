package service

import (
	"context"
	"testing"

	order "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
)

func TestDeleteOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteOrderService(ctx)
	// init req and assert value

	req := &order.DeleteOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
