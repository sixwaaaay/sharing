package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

// FavoriteActionHandler 生成点赞操作 handler
func FavoriteActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FavoriteReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FavoriteResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		favoriteLogic := logic.NewFavoriteLogic(c.Request.Context(), appCtx)
		resp, err := favoriteLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FavoriteResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
