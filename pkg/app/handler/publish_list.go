package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

func PublishListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.PubListReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.PubListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		publishListLogic := logic.NewPublishListLogic(c.Request.Context(), appCtx)
		resp, err := publishListLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.PubListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
