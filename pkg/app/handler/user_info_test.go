package handler

import (
	"context"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/common/testhelper"
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestUserInfo(t *testing.T) {
	const path = "/douyin/user/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "user info,biz logic success", // 业务逻辑成功
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"12345"},
				"token":   {"token"},
			},
			Expected: `{"status_code":0,"status_msg":"user info success","user":{"follow_count":100, "follower_count":101, "id":23, "is_follow":false, "name":"name"}}`,
		},
		{
			Name:   "user info,biz logic fail", // 业务逻辑失败，返回错误信息
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"987"},
				"token":   {"token"},
			},
			Expected: `{"status_code":1001,"status_msg":"user not found"}`,
		},
		{
			Name:     "user info, param error", // 参数错误
			Method:   "GET",
			Path:     path,
			Form:     url.Values{},
			Expected: `{"status_code":1001,"status_msg":"invalid params"}`,
		},
	}

	// 替换依赖的业务逻辑
	// 用于验证 handler 层的正确和错误分支
	logic.NewUserInfoLogic = func(ctx context.Context, appCtx *service.AppContext) logic.UserInfoLogic {
		return func(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
			if req.UserID == 987 {
				return nil, errorx.NewDefaultCodeErr("user not found")
			}
			msg := "user info success"
			return &types.UserInfoResp{
				StatusCode: 0,
				StatusMsg:  &msg,
				User: &types.User{
					FollowCount:   100,
					FollowerCount: 101,
					ID:            23,
					IsFollow:      false,
					Name:          "name",
				},
			}, nil
		}
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name+" "+testCase.Method+" "+testCase.Path, func(t *testing.T) {
			w, _ := testhelper.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
		})
	}
}
