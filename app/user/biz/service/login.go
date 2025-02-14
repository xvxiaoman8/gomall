package service

import (
	"context"
	"errors"

	"github.com/xvxiaoman8/gomall/app/user/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/user/biz/dal/redis"
	"github.com/xvxiaoman8/gomall/app/user/biz/model"
	user "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	// 参数校验
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	// 获取用户信息
	row, err := model.GetByEmail(s.ctx, mysql.DB, redis.RedisClient, req.Email)
	if err != nil {
		return nil, err
	}
	// 密码比对
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	// 组织响应
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return resp, nil
}
