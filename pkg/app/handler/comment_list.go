package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

func CommentListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CommentListReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.CommentListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		commentListLogic := logic.NewCommentListLogic(c.Request.Context(), appCtx)
		resp, err := commentListLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.CommentListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
