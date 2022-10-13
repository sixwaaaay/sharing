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

func TestFollowedListHandler(t *testing.T) {
	logic.NewFollowedListLogic = func(ctx context.Context, appCtx *service.AppContext) logic.FollowedListLogic {
		return func(req *types.RelationReq) (*types.FollowListResp, error) {
			if req.UserId == 101 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}
			resp := &types.FollowListResp{
				StatusCode: 0,
				StatusMsg:  nil,
				UserList:   nil,
			}
			return resp, nil
		}
	}
	const path = "/douyin/relation/follow/list/"
	var testCase = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试获取关注列表，业务逻辑成功
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"100"},
				"token":   {"token"},
			},
			Expected: `{"status_code":0,"status_msg":null,"user_list":null}`,
		},
		{
			Name:   "biz logic fail", // 测试获取关注列表，业务逻辑失败, 返回错误信息
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"101"},
				"token":   {"token"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error","user_list":null}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"douyin"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","user_list":null}`,
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
