package handler

import (
	"bytelite/app/favorite/logic"
	"bytelite/app/favorite/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
