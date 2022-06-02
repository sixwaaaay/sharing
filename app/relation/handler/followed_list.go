package handler

import (
	"bytelite/app/relation/logic"
	"bytelite/app/relation/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func FollowedListHandler(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.RelationReq
		if err := c.ShouldBind(&req); err != nil {
			codeError := errorx.NewDefaultCodeErr("invalid params")
			resp := &types.FollowListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
			return
		}

		followedListLogic := logic.NewFollowedListLogic(c.Request.Context(), appCtx)
		resp, err := followedListLogic(&req)
		if err != nil {
			codeError := err.(*errorx.CodeError)
			resp = &types.FollowListResp{}
			copier.Copy(resp, codeError)
			c.JSON(200, resp)
		} else {
			c.JSON(200, resp)
		}
	}
}
