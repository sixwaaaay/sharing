package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func FollowedListHandler(ctx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.RelationReq, types.FollowListResp](ctx, func(ctx context.Context, context *service.AppContext) func(*types.RelationReq) (*types.FollowListResp, error) {
		return logic.NewFollowedListLogic(ctx, context)
	})
}
