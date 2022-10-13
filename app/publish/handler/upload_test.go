package handler

import (
	"bytelite/app/publish/logic"
	"bytelite/app/publish/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"bytes"
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestBinding(t *testing.T) {
	// 测试传输视频时数据的绑定
	videoFile := MockFile{
		Filename: "video.mp4",
		Content:  []byte("video content"), // 模拟上传的文件内容
	}
	testToken := "testToken"
	testTitle := "testTitle"
	u := "/douyin/publish/action/"
	method := "POST"
	req := createPublishReq(t, &testToken, &testTitle, videoFile, u, method)
	var actual types.UploadReq
	// 解析请求参数
	err := binding.FormMultipart.Bind(req, &actual)
	assert.NoError(t, err)
	// 检查参数是否绑定成功
	assert.Equal(t, testToken, actual.Token)
	assert.Equal(t, testTitle, actual.Title)
	assert.Equal(t, videoFile.Filename, actual.File.Filename)
	open, err := actual.File.Open()
	assert.NoError(t, err)
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			t.Error(err)
		}
	}(open)
	// 检查内容
	content, err := io.ReadAll(open)
	assert.NoError(t, err)
	assert.Equal(t, videoFile.Content, content)
}

type PublishTestCase struct {
	Name     string
	Method   string
	Path     string
	Token    *string
	Title    *string
	File     MockFile
	Expected string
}

func TestUploadHandler(t *testing.T) {
	logic.NewUploadLogic = func(ctx context.Context, appCtx *service.AppContext) logic.UploadLogic {
		return func(req *types.UploadReq) (*types.UploadResp, error) {
			if req.Title == "error" {
				return nil, errorx.NewDefaultError("error")
			}
			return &types.UploadResp{
				StatusCode: 0,
				StatusMsg:  nil,
			}, nil
		}
	}
	token := "testToken"
	title := "testTitle"
	errorTitle := "error"
	const path = "/douyin/publish/action/"
	mockFile := MockFile{
		Filename: "video.mp4",
		Content:  []byte("video content"),
	} // 模拟上传的文件内容
	cases := []PublishTestCase{
		{
			Name:     "success", // 业务逻辑正常
			Method:   "POST",
			Path:     path,
			Token:    &token,
			Title:    &title,
			File:     mockFile,
			Expected: `{"status_code":0,"status_msg":null}`,
		},
		{
			Name:     "logic error", // 业务逻辑错误，返回错误信息
			Method:   "POST",
			Path:     path,
			Token:    &token,
			Title:    &errorTitle,
			File:     mockFile,
			Expected: `{"status_code":1001,"status_msg":"error"}`,
		},
		{
			Name:     "param error", // 参数错误
			Method:   "POST",
			Path:     path,
			Token:    nil,
			Title:    &title,
			File:     mockFile,
			Expected: `{"status_code":1001,"status_msg":"invalid params"}`,
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := createPublishReq(t, c.Token, c.Title, c.File, c.Path, c.Method)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, c.Expected, w.Body.String())
		})
	}
}

type MockFile struct {
	Filename string
	Content  []byte
}

func createPublishReq(t *testing.T, token *string, title *string, file MockFile, u string, method string) *http.Request {
	var body bytes.Buffer

	// 写入测试文件
	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile("data", file.Filename)
	assert.NoError(t, err)

	n, err := fw.Write(file.Content)
	assert.NoError(t, err)
	assert.Equal(t, len(file.Content), n)

	// 写入token字段
	if token != nil {
		field, err := mw.CreateFormField("token")
		assert.NoError(t, err)
		n, err = field.Write([]byte(*token))
		assert.NoError(t, err)
		assert.Equal(t, len(*token), n)
	}

	// 写入title字段
	if title != nil {
		field, err := mw.CreateFormField("title")
		assert.NoError(t, err)
		n, err = field.Write([]byte(*title))
		assert.NoError(t, err)
		assert.Equal(t, len(*title), n)
		err = mw.Close()
		assert.NoError(t, err)
	}

	req, err := http.NewRequest(method, u, &body)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "multipart/form-data"+"; boundary="+mw.Boundary())
	return req
}
