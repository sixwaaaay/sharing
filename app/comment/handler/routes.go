package handler

import (
	"bytelite/service"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/comment/action/", CommentActionHandler(appCtx))
	server.GET("/comment/list/", CommentListHandler(appCtx))
}
