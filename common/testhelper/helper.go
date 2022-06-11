package testhelper

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// GenRequest mock 一个http请求并生成一个httptest.ResponseRecorder
func GenRequest(r *gin.Engine, method, u string, body io.Reader, form url.Values) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, u, body)
	c.Request.PostForm = form
	r.ServeHTTP(w, c.Request)
	return w, c
}
