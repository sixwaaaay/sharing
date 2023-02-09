package handler

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
)

func FollowActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.FollowActionReq, types.FollowActionResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.FollowActionReq) (*types.FollowActionResp, error) {
		return logic.NewFollowActionLogic(ctx, context)
	})
}
