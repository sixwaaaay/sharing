package handler

import (
	"github.com/gin-gonic/gin"
	"testing"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	group := r.Group("douyin")
	RegisterHandlers(group, nil)
	m.Run()
}
