package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

func FollowActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FollowActionReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FollowActionResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		followActionLogic := logic.NewFollowActionLogic(c.Request.Context(), appCtx)
		resp, err := followActionLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FollowActionResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
