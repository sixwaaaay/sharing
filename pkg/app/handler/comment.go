package handler

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
)

func WrapHandler[Req any, Resp any](appContext *service.AppContext, builder func(context.Context, *service.AppContext) func(*Req) (*Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Req)
		if err := c.ShouldBind(req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := new(Resp)
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}
		handler := builder(c.Request.Context(), appContext)
		resp, err := handler(req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = new(Resp)
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}

func CommentActionHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return WrapHandler[types.CommentReq, types.CommentResp](appCtx, func(ctx context.Context, appCtx *service.AppContext) func(*types.CommentReq) (*types.CommentResp, error) {
		return logic.NewCommentLogic(ctx, appCtx)
	})
}
