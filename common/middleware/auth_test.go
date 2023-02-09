package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sixwaaaay/sharing/common/auth"
	"github.com/sixwaaaay/sharing/common/testhelper"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	context := service.AppContext{
		JWTSigner: auth.NewJWTSigner("secret"),
	}
	r.Use(VerifyToken(&context))
	var testUserId int64 = 25
	token, err := context.JWTSigner.GenerateToken(testUserId, 60*60*24*7)
	require.NoError(t, err)

	r.GET("/test", func(c *gin.Context) {
		ctx := c.Request.Context()
		value := ctx.Value(UserClaimsKey).(int64)
		require.Equal(t, testUserId, value)
		c.String(200, "test")
	})
	w, _ := testhelper.GenRequest(r, "GET", "/test", nil, url.Values{
		"token": {token},
	})
	require.Equal(t, w.Code, 200)
}
