package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"context"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/common/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
	"testing"
)

func TestFavoriteActionHandler(t *testing.T) {
	// 模拟逻辑层
	logic.NewFavoriteLogic = func(ctx context.Context, appCtx *service.AppContext) logic.FavoriteActionLogic {
		return func(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
			logx.WithContext(ctx).Infof("req: %+v", req)
			if req.UserId == 101 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}
			resp = &types.FavoriteResp{
				StatusCode: 0,
				StatusMsg:  "success",
			}
			return resp, nil
		}
	}

	const path = "/douyin/favorite/action/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试点赞，业务逻辑成功
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id":     {"100"},
				"token":       {"token"},
				"video_id":    {"100"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":0,"status_msg":"success"}`,
		},
		{
			Name:   "biz logic fail", // 测试点赞，业务逻辑失败, 返回错误信息
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id":     {"101"},
				"token":       {"token"},
				"video_id":    {"100"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error"}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id":     {"douyin"}, // 非数字
				"token":       {"token"},
				"video_id":    {"100"},
				"action_type": {"1"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params"}`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name+" "+testCase.Method+" "+testCase.Path, func(t *testing.T) {
			w, _ := testhelper.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
			t.Logf("%s", w.Body.String())
		})
	}
}
