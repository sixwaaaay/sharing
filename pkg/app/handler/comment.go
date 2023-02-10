package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func CommentActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.CommentReq, types.CommentResp](appCtx, func(ctx context.Context, appCtx *service.AppContext) func(*types.CommentReq) (*types.CommentResp, error) {
		return logic.NewCommentLogic(ctx, appCtx)
	})
}

func NewCommentActionHandler(appCtx *service.AppContext) *Handler {
	return &Handler{
		Handler: CommentActionHandler(appCtx),
		Path:    "/comment/action/",
		Method:  "POST",
	}
}
