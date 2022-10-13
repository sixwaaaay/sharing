package handler

import (
	"bytelite/app/relation/logic"
	"bytelite/app/relation/types"
	"bytelite/common/errorx"
	"bytelite/common/testhelper"
	"bytelite/service"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFollowActionHandler(t *testing.T) {
	logic.NewFollowActionLogic = func(ctx context.Context, appCtx *service.AppContext) logic.FollowLogic {
		return func(req *types.FollowActionReq) (*types.FollowActionResp, error) {
			if req.ToUserId == 101 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}
			resp := &types.FollowActionResp{
				StatusCode: 0,
				StatusMsg:  "ok",
			}
			return resp, nil
		}
	}
	const path = "/douyin/relation/action/"
	var testCase = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试关注，业务逻辑成功
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"token":       {"token"},
				"to_user_id":  {"100"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":0,"status_msg":"ok"}`,
		},
		{
			Name:   "biz logic fail", // 测试关注，业务逻辑失败, 返回错误信息
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"token":       {"token"},
				"to_user_id":  {"101"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error"}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"token":       {"token"},
				"to_user_id":  {"douyin"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params"}`,
		},
	}
	for _, testCase := range testCase {
		t.Run(testCase.Name+" "+testCase.Method+" "+testCase.Path, func(t *testing.T) {
			w, _ := testhelper.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
			t.Logf("%s", w.Body.String())
		})
	}
}
