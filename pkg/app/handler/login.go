package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// Login 生成用户注册的handler
func Login(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.UserReq, types.UserResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.UserReq) (*types.UserResp, error) {
		return logic.NewLoginLogic(ctx, context)
	})
}
