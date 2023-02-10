package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func PublishListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.PubListReq, types.PubListResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.PubListReq) (*types.PubListResp, error) {
		return logic.NewPublishListLogic(ctx, context)
	})
}

func NewPublishListHandler(appCtx *service.AppContext) *Handler {
	return &Handler{
		Handler: PublishListHandler(appCtx),
		Path:    "/publish/list/",
		Method:  "GET",
	}
}
