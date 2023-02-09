package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
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
