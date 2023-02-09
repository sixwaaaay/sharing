package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/configs"
	"github.com/sixwaaaay/sharing/pkg/app/handler"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/common/middleware"
	"github.com/zeromicro/go-zero/core/conf"
	"go.uber.org/zap"
	"time"
)

func main() {
	var c configs.Config
	conf.MustLoad("configs/config.yaml", &c)
	r := gin.New()

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	r.Use(ginzap.RecoveryWithZap(logger, true))

	appCtx := service.NewAppContext(&c)
	group := r.Group("/douyin")
	group.POST("/user/register/", handler.Register(appCtx))
	group.POST("/user/login/", handler.Login(appCtx))
	group.Use(middleware.VerifyToken(appCtx)) // 中间件验证token
	authHook := middleware.Authority(appCtx)  // 中间件验证权限
	group.GET("/user/", authHook, handler.UserInfoHandler(appCtx))
	handler.RegisterFeedHandlers(group, appCtx)
	group.Use(authHook)
	handler.RegisterPublishHandlers(group, appCtx)
	handler.RegisterCommentHandlers(group, appCtx)
	handler.RegisterFavorHandlers(group, appCtx)
	handler.RegisterRelationHandlers(group, appCtx)
	err = r.Run(c.Addr())
	if err != nil {
		panic(err)
	}
}
