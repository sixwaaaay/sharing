package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

func CommentActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CommentReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.CommentResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		commentActionLogic := logic.NewCommentLogic(c.Request.Context(), appCtx)
		resp, err := commentActionLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.CommentResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
