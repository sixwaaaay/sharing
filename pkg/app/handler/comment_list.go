package handler

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
)

func CommentListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.CommentListReq, types.CommentListResp](appCtx, func(ctx context.Context, appCtx *service.AppContext) func(*types.CommentListReq) (*types.CommentListResp, error) {
		return logic.NewCommentListLogic(ctx, appCtx)
	})
}
