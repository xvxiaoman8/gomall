package service

import (
	"context"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

//func TestGetProduct_Run(t *testing.T) {
//	ctx := context.Background()
//	s := NewGetProductService(ctx)
//	// init req and assert value
//
//	req := &product.GetProductRequest{}
//	resp, err := s.Run(req)
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//	if resp == nil {
//		t.Errorf("unexpected nil response")
//	}
//	// // todo: edit your unit test
//}
