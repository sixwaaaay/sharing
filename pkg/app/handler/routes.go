package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/service"
)

func RegisterCommentHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/comment/action/", CommentActionHandler(appCtx))
	server.GET("/comment/list/", CommentListHandler(appCtx))
}

func RegisterFavorHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/favorite/action/", FavoriteActionHandler(appCtx))
	server.GET("/favorite/list/", FavoriteListHandler(appCtx))
}

// RegisterFeedHandlers 将路由注册到gin的路由组中
func RegisterFeedHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.GET("/feed/", Feed(appCtx))
}

func RegisterPublishHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.GET("/publish/list/", PublishListHandler(appCtx))
	server.POST("/publish/action/", UploadHandler(appCtx))
}

func RegisterRelationHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/relation/action/", FollowActionHandler(appCtx))
	server.GET("/relation/follow/list/", FollowedListHandler(appCtx))
	server.GET("/relation/follower/list/", FollowerListHandler(appCtx))
}

// RegisterUserHandlers 将路由注册到gin的路由组中
func RegisterUserHandlers(server *gin.RouterGroup, appCtx *service.AppContext) {
	server.POST("/user/register/", Register(appCtx))
	server.POST("/user/login/", Login(appCtx))
	server.GET("/user/", UserInfoHandler(appCtx))
}
