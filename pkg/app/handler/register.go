package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// Register 用户注册处理
func Register(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.UserReq, types.UserResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.UserReq) (*types.UserResp, error) {
		return logic.NewRegisterLogic(ctx, context)
	})
}
