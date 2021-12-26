package middleware

import (
	"github.com/go-errors/errors"
	"github.com/wa1kman999/cowman/internal/http/vs"
	"github.com/wa1kman999/cowman/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			vs.SendBadData(ctx, errors.New("未登录或非法访问"), nil)
			ctx.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				vs.SendBadData(ctx, errors.New("授权已过期"), nil)
				ctx.Abort()
				return
			}
			vs.SendBadData(ctx, err, nil)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
