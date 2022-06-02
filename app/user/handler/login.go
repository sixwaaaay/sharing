package handler

import (
	"bytelite/app/user/logic"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
