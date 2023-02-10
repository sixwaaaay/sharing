package main

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/configs"
	_ "github.com/sixwaaaay/sharing/docs"
	"github.com/sixwaaaay/sharing/pkg/app/handler"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/common/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeromicro/go-zero/core/conf"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// @title sharing
// @version 1.0
// @description a simple short-form, video-sharing app backend
// @contact.name sixwaaaay
// @contact.url  https://github.com/sixwaaaay
// @license.name Apache 2.0
// @license.url https://github.com/sixwaaaay/sharing/blob/master/LICENSE
// @host localhost:8080
func main() {
	fx.New(fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger}
	}), fx.Provide(
		NewConfig,
		service.NewAppContext,
		NewEngine,
		NewLogger),
		handler.HandlerMoudle,
		fx.Invoke(Register, NewServer)).Run()
}
func NewConfig() *configs.Config {
	c := new(configs.Config)
	conf.MustLoad("configs/config.yaml", c)
	return c
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewEngine(logger *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r
}

func NewServer(lc fx.Lifecycle, r *gin.Engine, c *configs.Config) {
	srv := &http.Server{Addr: c.Addr(), Handler: r}
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}()
		return nil
	}, OnStop: func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}})
}

type RouterOption struct {
	fx.In
	E      *gin.Engine
	AppCtx *service.AppContext
	Pubs   []*handler.Handler `group:"public"`
	Opt    []*handler.Handler `group:"option"`
	Pri    []*handler.Handler `group:"private"`
}

func Register(opt RouterOption) {
	opt.E.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r := opt.E.Group("/douyin")
	for _, h := range opt.Pubs {
		r.Handle(h.Method, h.Path, h.Handler)
	}
	r.Use(middleware.VerifyToken(opt.AppCtx))    // 中间件验证token
	authHook := middleware.Authority(opt.AppCtx) // 中间件验证权限
	for _, h := range opt.Opt {
		r.Handle(h.Method, h.Path, authHook, h.Handler)
	}
	r.Use(authHook)
	for _, h := range opt.Pri {
		r.Handle(h.Method, h.Path, authHook, h.Handler)
	}
}
