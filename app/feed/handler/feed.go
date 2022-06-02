package handler

import (
	"bytelite/app/feed/logic"
	"bytelite/app/feed/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// Feed 生成视频 Feed 流 handler
func Feed(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FeedReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FeedResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		feedLogic := logic.NewFeedLogic(c.Request.Context(), appCtx)
		resp, err := feedLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FeedResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
