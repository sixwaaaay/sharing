package handler

import (
	"bytelite/app/user/logic"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// UserInfoHandler 生成获取用户信息的handler
func UserInfoHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserInfoReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.UserInfoResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}
		logicHandle := logic.NewUserInfoLogic(c.Request.Context(), appCtx)
		resp, err := logicHandle(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.UserInfoResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
