package main

import (
	"context"
	payment "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"github.com/xvxiaoman8/gomall/app/payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}

// Refund implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Refund(ctx context.Context, req *payment.RefundReq) (resp *payment.RefundResp, err error) {
	resp, err = service.NewRefundService(ctx).Run(req)

	return resp, err
}
