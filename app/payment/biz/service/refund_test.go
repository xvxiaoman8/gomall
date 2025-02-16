package service

import (
	"context"
	payment "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestRefund_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefundService(ctx)
	// init req and assert value

	req := &payment.RefundReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
