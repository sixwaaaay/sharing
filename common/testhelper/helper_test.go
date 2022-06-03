package testhelper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestGenRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	type CustomParams struct {
		Name string `json:"name" form:"name"`
	}

	r.GET("/test", func(c *gin.Context) {
		var params CustomParams
		_ = c.ShouldBind(&params)
		c.JSON(200, &params)
	})
	form := url.Values{
		"name": {"test"},
	}
	expected := `{"name":"test"}`
	request, _ := GenRequest(r, "GET", "/test", nil, form)
	require.JSONEq(t, expected, request.Body.String())
	require.Equal(t, 200, request.Code)
}
