package service

import (
	"context"
	"testing"
	user "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/user"
)

func TestModify_Run(t *testing.T) {
	ctx := context.Background()
	s := NewModifyService(ctx)
	// init req and assert value

	req := &user.ModifyReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
