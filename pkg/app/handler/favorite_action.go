package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// FavoriteActionHandler 生成点赞操作 handler
func FavoriteActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.FavoriteReq, types.FavoriteResp](appCtx, func(ctx context.Context, appCtx *service.AppContext) func(*types.FavoriteReq) (*types.FavoriteResp, error) {
		return logic.NewFavoriteLogic(ctx, appCtx)
	})
}
