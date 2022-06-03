package handler

import (
	"bytelite/service"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.GET("/publish/list/", PublishListHandler(appCtx))
	server.POST("/publish/action/", UploadHandler(appCtx))
}
