package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

// FavoriteListHandler 生成获取点赞列表的handler
func FavoriteListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FavoriteListReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FavoriteListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		favoriteListLogic := logic.NewFavoriteListLogic(c.Request.Context(), appCtx)
		resp, err := favoriteListLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FavoriteListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
