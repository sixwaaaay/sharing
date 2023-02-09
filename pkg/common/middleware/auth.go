package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/zeromicro/go-zero/core/logx"
)

const UserClaimsKey = "user"

type Auth struct {
	Token *string `form:"token" json:"token" binding:"required"`
}

// VerifyToken 将 token 转化为 user_id 如果可行的话
func VerifyToken(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var auth Auth
		if err := c.ShouldBind(&auth); err != nil {
			c.Next()
		}
		if auth.Token == nil {
			c.Next()
			return
		}
		userId, err := appCtx.JWTSigner.ValidateToken(*auth.Token)
		if err != nil {
			logx.Error(err)
			c.AbortWithStatus(401)
			return
		}
		logx.Infof("userId: %d", userId)
		ctx := context.WithValue(c.Request.Context(), UserClaimsKey, userId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Authority 过滤请求
func Authority(appCtx *service.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Request.Context().Value(UserClaimsKey).(int64)
		if !ok {
			c.AbortWithStatus(401)
			return
		}
	}
}
