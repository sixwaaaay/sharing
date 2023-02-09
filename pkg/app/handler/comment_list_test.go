package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	testhelper2 "github.com/sixwaaaay/sharing/pkg/common/testhelper"

	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestCommentListHandler(t *testing.T) {
	logic.NewCommentListLogic = func(ctx context.Context, appCtx *service.AppContext) logic.CommentListLogic {
		return func(req *types.CommentListReq) (*types.CommentListResp, error) {
			if req.VideoId == 321 {
				return nil, errorx.NewDefaultError("test error")
			}
			return &types.CommentListResp{
				StatusCode: 0,
				StatusMsg:  nil,
				CommentList: []types.Comment{
					{
						Content:    "test comment",
						CreateDate: "06-20",
						ID:         1,
						User: types.User{
							FollowCount:   23,
							FollowerCount: 23,
							ID:            1,
							IsFollow:      false,
							Name:          "a user",
						},
					},
				},
			}, nil
		}
	}

	const path = "/douyin/comment/list/"
	var testCases = []testhelper2.TestCase{
		{
			Name:   "biz logic success", // 测试获取评论列表，业务逻辑成功
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"video_id": {"100"},
				"token":    {"token"},
			},
			Expected: `{"status_code":0,"status_msg":null,"comment_list":[{"content":"test comment","create_date":"06-20","id":1,"user":{"follow_count":23,"follower_count":23,"id":1,"is_follow":false,"name":"a user"}}]}`,
		},
		{
			Name:   "biz logic fail", // 测试获取评论列表，业务逻辑失败, 返回错误信息
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"video_id": {"321"},
				"token":    {"token"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error","comment_list":null}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"video_id": {"douyin"}, //非 ID 正确的数据类型
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","comment_list":null}`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name+" "+testCase.Method+" "+testCase.Path, func(t *testing.T) {
			w, _ := testhelper2.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
			t.Logf("%s", w.Body.String())
		})
	}

}
