package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// UserInfoHandler 生成获取用户信息的 handler
func UserInfoHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.UserInfoReq, types.UserInfoResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.UserInfoReq) (*types.UserInfoResp, error) {
		return logic.NewUserInfoLogic(ctx, context)
	})
}

func NewUserInfoHandler(appCtx *service.AppContext) *Handler {
	return &Handler{
		Handler: UserInfoHandler(appCtx),
		Path:    "/user/info/",
		Method:  "GET",
	}
}
