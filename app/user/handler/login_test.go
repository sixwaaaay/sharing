package handler

import (
	"bytelite/app/user/logic"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/common/testhelper"
	"bytelite/service"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestLogin(t *testing.T) {
	const path = "/douyin/user/login/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "login,biz logic success", // 测试登录，业务逻辑成功
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"test"},
				"password": {"123456"},
			},
			Expected: `{"status_code":0,"status_msg":"success","token":"token for test success","user_id":47777}`,
		},
		{
			Name:   "login,biz logic fail", // 测试登录，业务逻辑失败, 返回错误信息
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"fail"},
				"password": {"fail"},
			},
			Expected: `{"status_code":1001,"status_msg":"fail to login","token":"","user_id":0}`,
		},
		{
			Name:   "login,params error", // 参数错误
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"test"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","token":"","user_id":0}`,
		},
	}
	// 替换依赖的业务逻辑
	logic.NewLoginLogic = loginHookFunc
	for _, testCase := range testCases {
		t.Run(testCase.Name+" "+testCase.Method+" "+testCase.Path, func(t *testing.T) {
			w, _ := testhelper.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
		})
	}
}

/*
替换依赖的业务逻辑
用于验证 handler 层的正确和错误分支
*/
func loginHookFunc(_ context.Context, _ *service.AppContext) logic.LoginLogic {
	return func(req *types.UserReq) (resp *types.UserResp, err error) {
		if req.Username == "fail" && req.Password == "fail" {
			return nil, errorx.NewDefaultCodeErr("fail to login")
		}
		return &types.UserResp{
			StatusCode: 0,
			StatusMsg:  "success",
			Token:      "token for test success",
			UserID:     47777,
		}, nil
	}
}
