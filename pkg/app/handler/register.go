package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
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
