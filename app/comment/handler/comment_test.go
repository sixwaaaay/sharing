package handler

import (
	"bytelite/app/comment/logic"
	"bytelite/app/comment/types"
	"bytelite/common/cotypes"
	"bytelite/common/errorx"
	"bytelite/common/testhelper"
	"bytelite/service"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
	"testing"
)

func TestCommentActionHandler(t *testing.T) {
	logic.NewCommentLogic = func(ctx context.Context, appCtx *service.AppContext) logic.CommentLogic {
		return func(req *types.CommentReq) (*types.CommentResp, error) {
			logx.WithContext(ctx).Infof("req: %+v", req)
			if req.UserId == 255 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}
			resp := &types.CommentResp{
				StatusCode: 0,
				StatusMsg:  nil,
				Comment: &cotypes.Comment{
					Content:    "comment content",
					CreateDate: "11-26",
					ID:         123,
					User: cotypes.User{
						FollowCount:   2,
						FollowerCount: 3,
						ID:            456,
						IsFollow:      false,
						Name:          "test name",
					},
				},
			}
			return resp, nil
		}
	}
	const path = "/douyin/comment/action/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试评论，业务逻辑成功
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id":      {"100"},
				"token":        {"token"},
				"video_id":     {"123"},
				"action_type":  {"1"},
				"comment_text": {"comment content"},
			},
			Expected: `{"status_code":0,"status_msg":null,"comment":{"content":"comment content","create_date":"11-26","id":123,"user":{"follow_count":2,"follower_count":3,"id":456,"is_follow":false,"name":"test name"}}}`,
		},
		{
			Name:   "biz logic fail", // 测试评论，业务逻辑失败, 返回错误信息
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id":  {"255"},
				"token":    {"token"},
				"video_id": {"123"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error","comment":null}`,
		},
		{
			Name:   "params error", // 参数错误
			Method: "POST",
			Path:   path,
			Form: url.Values{
				"user_id": {"douyin"},
			},
			Expected: `{"status_code":1001,"status_msg":"invalid params","comment":null}`,
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
