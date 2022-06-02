package handler

import (
	"bytelite/service"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers 将路由注册到gin的路由组中
func RegisterHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/user/register/", Register(appCtx))
	server.POST("/user/login/", Login(appCtx))
	server.GET("/user/", UserInfoHandler(appCtx))
}
