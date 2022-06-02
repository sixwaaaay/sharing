package handler

import (
	"bytelite/service"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/favorite/action/", FavoriteActionHandler(appCtx))
	server.GET("/favorite/list/", FavoriteListHandler(appCtx))
}
