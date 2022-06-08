package main

import (
	comment "bytelite/app/comment/handler"
	favorite "bytelite/app/favorite/handler"
	feed "bytelite/app/feed/handler"
	publish "bytelite/app/publish/handler"
	relation "bytelite/app/relation/handler"
	user "bytelite/app/user/handler"
	"bytelite/common/middleware"
	"bytelite/etc"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var c etc.Config
	conf.MustLoad("etc/config.yaml", &c)
	r := gin.Default()
	appCtx := service.NewAppContext(&c)
	group := r.Group("/douyin")
	group.POST("/user/register/", user.Register(appCtx))
	group.POST("/user/login/", user.Login(appCtx))
	group.Use(middleware.VerifyToken(appCtx))
	authHook := middleware.Authority(appCtx)
	group.GET("/user/", authHook, user.UserInfoHandler(appCtx))
	feed.RegisterHandlers(group, appCtx)
	group.Use(authHook)
	publish.RegisterHandlers(group, appCtx)
	comment.RegisterHandlers(group, appCtx)
	favorite.RegisterHandlers(group, appCtx)
	relation.Register(group, appCtx)
	err := r.Run(c.Addr())
	if err != nil {
		panic(err)
	}
}
