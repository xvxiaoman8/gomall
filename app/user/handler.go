package main

import (
	"context"
	"github.com/xvxiaoman8/gomall/app/user/biz/service"
	user "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Modify implements the UserServiceImpl interface.
func (s *UserServiceImpl) Modify(ctx context.Context, req *user.ModifyReq) (resp *user.ModifyResp, err error) {
	resp, err = service.NewModifyService(ctx).Run(req)

	return resp, err
}
