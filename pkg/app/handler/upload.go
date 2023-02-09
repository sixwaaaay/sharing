package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func UploadHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.UploadReq, types.UploadResp](appCtx, func(ctx context.Context, context *service.AppContext) func(*types.UploadReq) (*types.UploadResp, error) {
		return logic.NewUploadLogic(ctx, context)
	})
}
