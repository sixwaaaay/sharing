package handler

import (
	"bytelite/app/publish/logic"
	"bytelite/app/publish/types"
	"bytelite/common/cotypes"
	"bytelite/common/errorx"
	"bytelite/common/testhelper"
	"bytelite/service"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPublishListHandler(t *testing.T) {
	logic.NewPublishListLogic = func(ctx context.Context, appCtx *service.AppContext) logic.PubListLogic {
		return func(req *types.PubListReq) (*types.PubListResp, error) {
			if req.UserId == 101 {
				//raise error
				return nil, errorx.NewDefaultError("test error")
			}

			resp := &types.PubListResp{
				StatusCode: 0,
				StatusMsg:  nil,
				VideoList: []cotypes.Video{
					{
						Author: cotypes.User{
							FollowCount:   20,
							FollowerCount: 35,
							ID:            1110,
							IsFollow:      false,
							Name:          "test Case",
						},
						CommentCount:  10,
						CoverURL:      "http://test.com/test.jpg",
						FavoriteCount: 20,
						ID:            1235,
						IsFavorite:    false,
						PlayURL:       "http://test.com/test.mp4",
						Title:         "test video title",
					},
				},
			}
			return resp, nil
		}
	}
	const path = "/douyin/publish/list/"
	var testCases = []testhelper.TestCase{
		{
			Name:   "biz logic success", // 测试获取视频列表，业务逻辑成功
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"100"},
				"token":   {"token"},
			},
			Expected: `{
				"status_code":0,"status_msg":null,
				"video_list":[{
					"author":{"follow_count":20,"follower_count":35,"id":1110,"is_follow":false,"name":"test Case"},
					"comment_count":10,
					"cover_url":"http://test.com/test.jpg",
					"favorite_count":20,
					"id":1235,
					"is_favorite":false,
					"play_url":"http://test.com/test.mp4",
					"title":"test video title"}
				]}`},
		{
			Name:   "biz logic fail", // 测试获取视频列表，业务逻辑失败, 返回错误信息
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"101"},
				"token":   {"token"},
			},
			Expected: `{"status_code":1001,"status_msg":"test error","video_list":null}`,
		}, {
			Name:   "params error", // 参数错误
			Method: "GET",
			Path:   path,
			Form: url.Values{
				"user_id": {"douyin"}, //类型错误
				"token":   {"token"},
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
