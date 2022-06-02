package handler

import (
	"bytelite/app/relation/logic"
	"bytelite/app/relation/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func FollowerListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.RelationReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FollowerListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		followerListLogic := logic.NewFollowerListLogic(c.Request.Context(), appCtx)
		resp, err := followerListLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FollowerListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
