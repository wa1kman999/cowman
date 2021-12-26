package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "获取用户列表成功")
}
