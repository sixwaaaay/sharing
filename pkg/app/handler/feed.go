package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// Feed 生成视频 Feed 流 handler
func Feed(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.FeedReq, types.FeedResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.FeedReq) (*types.FeedResp, error) {
		return logic.NewFeedLogic(ctx, context)
	})
}
