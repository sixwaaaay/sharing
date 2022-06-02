package handler

import (
	"bytelite/app/user/logic"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// Register 用户注册处理
func Register(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.UserResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}
		logicHandle := logic.NewRegisterLogic(c.Request.Context(), appCtx)
		resp, err := logicHandle(&req)
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
