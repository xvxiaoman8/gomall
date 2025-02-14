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

type ModifyService struct {
	ctx context.Context
} // NewModifyService new ModifyService
func NewModifyService(ctx context.Context) *ModifyService {
	return &ModifyService{ctx: ctx}
}

// Run create note info
func (s *ModifyService) Run(req *user.ModifyReq) (resp *user.ModifyResp, err error) {
	// Finish your business logic.
	// 参数校验
	if req.Email == "" || req.OldPassword == "" || req.NewPassword == "" {
		return nil, errors.New("email or password is empty")
	}
	// 获取用户信息
	row, err := model.GetByEmail(s.ctx, mysql.DB, redis.RedisClient, req.Email)
	if err != nil {
		return nil, err
	}
	// 密码比对
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.OldPassword))
	if err != nil {
		return nil, err
	}
	// 对密码进行哈希加密
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 创建用户信息
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Modify(s.ctx, mysql.DB, redis.RedisClient, newUser)
	return
}
