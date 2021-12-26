package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userController "github.com/wa1kman999/cowman/internal/controller/user"
	"github.com/wa1kman999/cowman/internal/http/middleware"
)

const (
	v1prefix = "/cowman/v1"
)

// initRouter 初始化路由
func initRouter(router *gin.Engine) error {

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	user := router.Group(v1prefix + "/")
	// 登陆页面
	user.POST("login", userController.Login)
	//使用中间件
	userPrivate := user.Use(middleware.JWTAuth())
	userPrivate.GET("list", userController.GetList)

	return nil
}
