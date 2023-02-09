package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func FollowerListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.RelationReq, types.FollowerListResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.RelationReq) (*types.FollowerListResp, error) {
		return logic.NewFollowerListLogic(ctx, context)
	})
}
