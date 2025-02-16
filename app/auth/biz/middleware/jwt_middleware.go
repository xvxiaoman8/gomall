package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {

		// 获取 JWT Token
		tokenStr := ctx.Value("token").(string)
		if tokenStr == "" {
			return errors.New("token required")
		}

		// 解析并验证 Token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})
		if err != nil || !token.Valid {
			return errors.New("invalid token")
		}

		// 将 claims 存入上下文
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx = context.WithValue(ctx, "user", claims["user"])
			ctx = context.WithValue(ctx, "role", claims["role"])
		}

		return next(ctx, req, resp)
	}
}
