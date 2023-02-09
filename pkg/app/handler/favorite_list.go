package handler

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
)

// FavoriteListHandler 生成获取点赞列表的handler
func FavoriteListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.FavoriteListReq, types.FavoriteListResp](appCtx, func(ctx context.Context, appCtx *service.AppContext) func(*types.FavoriteListReq) (*types.FavoriteListResp, error) {
		return logic.NewFavoriteListLogic(ctx, appCtx)
	})
}
