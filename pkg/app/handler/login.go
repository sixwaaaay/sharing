package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

// Login 生成用户注册的handler
func Login(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.UserResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}
		loginLogic := logic.NewLoginLogic(c.Request.Context(), appCtx)
		resp, err := loginLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.UserResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
