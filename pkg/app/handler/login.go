package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// Login 生成用户注册的handler
// @Summary 用户登陆
// @Description 用户登陆
// @Tags 用户
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData types.UserReq true "用户信息"
// @Success 200 {object} types.UserResp
// @Router /douyin/user/login/ [post]
func Login(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.UserReq, types.UserResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.UserReq) (*types.UserResp, error) {
		return logic.NewLoginLogic(ctx, context)
	})
}

func NewLogin(appCtx *service.AppContext) *Handler {
	return &Handler{
		Handler: Login(appCtx),
		Path:    "/user/login/",
		Method:  "POST",
	}
}
