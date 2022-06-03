package handler

import (
	"github.com/gin-gonic/gin"
	"testing"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	RegisterHandlers(r.Group("/douyin"), nil)
	m.Run()
}
