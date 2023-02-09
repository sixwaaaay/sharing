package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/logic"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	testhelper2 "github.com/sixwaaaay/sharing/pkg/common/testhelper"

	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFeed(t *testing.T) {
	const path = "/douyin/feed/"
	var testCases = []testhelper2.TestCase{
		{
			Name:   "logic success", // 业务逻辑成功
			Method: "GET",
			Path:   path,
			Body:   nil,
			Form: url.Values{
				"latest_time": {"15897600"},
				"token":       {"token"},
			},
			Expected: `
	{"next_time":15897600,"status_code":0,"status_msg":null,
	 "video_list":
			[{"comment_count":123,"cover_url":"this is cover url",
       		  "favorite_count":322,"id":3413251,"is_favorite":false,
			"play_url":"this is video url",
			"title":"this is video title",
			"author":{"follow_count":3215,
			"follower_count":512523,"id":53253216216,
			"is_follow":true,"name":"username"}}]}`,
		},
		{
			Name:   "logic fail", //业务逻辑失败
			Method: "GET",
			Path:   path,
			Body:   nil,
			Form: url.Values{
				"latest_time": {"6"},
			},
			Expected: `{"status_code":1001,"status_msg":"time is too small","video_list":null,"next_time":null}`,
		},
		{
			Name:   "param fail", // 参数异常
			Method: "GET",
			Path:   path,
			Body:   nil,
			Form: url.Values{
				"latest_time": {"aaa"}, // 非数字
			},
			Expected: `{"status_code":1001, "status_msg":"invalid params","video_list":null,"next_time":null}`,
		},
	}
	// 模拟逻辑层
	logic.NewFeedLogic = func(ctx context.Context, appCtx *service.AppContext) logic.FeedLogic {
		return func(req *types.FeedReq) (*types.FeedResp, error) {
			var earliest int64 = 15897600
			// 模拟逻辑异常
			if *req.LatestTime < 500 {
				return nil, errorx.NewDefaultError("time is too small")
			}
			return &types.FeedResp{
				NextTime:   &earliest,
				StatusCode: 0,
				StatusMsg:  nil,
				VideoList: []types.Video{
					{
						CommentCount:  123,
						CoverURL:      "this is cover url",
						FavoriteCount: 322,
						ID:            3413251,
						IsFavorite:    false,
						PlayURL:       "this is video url",
						Title:         "this is video title",
						Author: types.User{
							FollowCount:   3215,
							FollowerCount: 512523,
							ID:            53253216216,
							IsFollow:      true,
							Name:          "username",
						},
					},
				},
			}, nil
		}
	}
	gin.SetMode(gin.TestMode)
	r := gin.New()
	group := r.Group("/douyin")
	RegisterFeedHandlers(group, nil)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			w, _ := testhelper2.GenRequest(r, testCase.Method, testCase.Path, testCase.Body, testCase.Form)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, testCase.Expected, w.Body.String())
		})
	}
}
