package handler

import (
	"bytelite/app/user/logic"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/common/testhelper"
	"bytelite/service"
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestRegister(t *testing.T) {
	const path = "/douyin/user/register/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "register, biz success", // 测试注册，业务逻辑成功
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"test"},
				"password": {"123456"},
			},
			Expected: `{"status_code":0,"status_msg":"success","token":"token for test success","user_id":47777}`,
		},
		{
			Name:   "register, biz error", // 测试注册，业务逻辑失败，返回错误信息
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"fail"},
				"password": {"fail"},
			},
			Expected: `{"status_code":1001,"status_msg":"fail to register","token":"","user_id":0}`,
		},
		{
			Name:   "register, param error", // 请求参数错误
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"username": {"test"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","token":"","user_id":0}`,
		},
	}
	// 替换依赖的业务逻辑
	logic.NewRegisterLogic = registerHookFunc
	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			w, _ := testhelper.GenRequest(r, c.Method, c.Path, nil, c.Form)
			require.Equal(t, http.StatusOK, w.Code)
			require.Equal(t, c.Expected, w.Body.String())
		})
	}
}

/*
替换依赖的业务逻辑
用于验证 handler 层的正确和错误分支
*/
func registerHookFunc(_ context.Context, _ *service.AppContext) logic.RegisterLogic {
	return func(req *types.UserReq) (resp *types.UserResp, err error) {
		if req.Username == "fail" && req.Password == "fail" {
			return nil, errorx.NewDefaultCodeErr("fail to register")
		}
		return &types.UserResp{
			StatusCode: 0,
			StatusMsg:  "success",
			Token:      "token for test success",
			UserID:     47777,
		}, nil
	}
}
