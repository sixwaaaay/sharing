package handler

import (
	"bytelite/service"
	"github.com/gin-gonic/gin"
)

func Register(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/relation/action/", FollowActionHandler(appCtx))
	server.GET("/relation/follow/list/", FollowedListHandler(appCtx))
	server.GET("/relation/follower/list/", FollowerListHandler(appCtx))
}
