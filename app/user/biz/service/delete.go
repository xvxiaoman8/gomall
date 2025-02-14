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

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
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
	// 创建用户信息
	User := &model.User{
		Email: req.Email,
	}
	err = model.Delete(s.ctx, mysql.DB, redis.RedisClient, User)
	return
}
