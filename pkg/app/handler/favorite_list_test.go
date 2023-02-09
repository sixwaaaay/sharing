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

func TestFavoriteListHandler(t *testing.T) {
	// 模拟逻辑层
	logic.NewFavoriteListLogic = func(ctx context.Context, appCtx *service.AppContext) logic.FavoriteListLogic {
		return func(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
			logx.WithContext(ctx).Infof("req: %+v", req)
			if req.UserId == 101 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}
			resp = &types.FavoriteListResp{
				StatusCode: 0,
				StatusMsg:  nil,
				VideoList:  nil,
			}
			return resp, nil
		}
	}

	const path = "/douyin/favorite/list/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试点赞列表，业务逻辑成功
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"100"},
				"token":   {"token"},
			},
			Expected: `{"status_code":0,"status_msg":null,"video_list":null}`,
		},
		{
			Name:   "biz logic fail", // 测试点赞列表，业务逻辑失败, 返回错误信息
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"101"},
				"token":   {"token"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error","video_list":null}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"douyin"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","video_list":null}`,
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
