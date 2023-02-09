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
	RegisterCommentHandlers(group, nil)
	RegisterUserHandlers(group, nil)
	RegisterRelationHandlers(group, nil)
	RegisterFavorHandlers(group, nil)
	RegisterFeedHandlers(group, nil)
	RegisterPublishHandlers(group, nil)
	m.Run()
}
