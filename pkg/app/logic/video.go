package logic

import (
	"context"
	"fmt"
	"github.com/sixwaaaay/sharing/common/covert"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/common/itertool"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"math"
)

// FetchField 填充相关字段
func FetchField(ctx context.Context, appCtx *service.AppContext, videoList []*dal.Videos, selfId int64) ([]types.Video, error) {
	multiUserInfo, err := QueryMultiUserInfo(ctx, appCtx, selfId, videoToUids(videoList))
	if err != nil {
		return nil, errorx.NewDefaultError("server error, query failed")
	}
	favorites, err := QueryVideoFavorites(ctx, appCtx, selfId, vIds(videoList))
	if err != nil {
		return nil, errorx.NewDefaultError("error to query data")
	}
	userMap := covert.UserMap(multiUserInfo)         // 用于join
	videos := toVideos(videoList)                    // 转换为所需类型
	favoriteMap := covert.Int64SliceToMap(favorites) // 点赞结果
	for i := 0; i < len(videoList); i++ {
		if user, ok := userMap[videoList[i].UserId]; ok {
			videos[i].Author = user
		}
		if _, ok := favoriteMap[videoList[i].Id]; ok {
			videos[i].IsFavorite = true
		}
		videos[i].PlayURL = fmt.Sprintf("%s%s", appCtx.ContentBaseUrl, videos[i].PlayURL)
		videos[i].CoverURL = fmt.Sprintf("%s%s", appCtx.ContentBaseUrl, videos[i].CoverURL)
	}
	return videos, nil
}

// QueryUserVideo 根据给定的用户id获取用户的视频列表
func QueryUserVideo(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) ([]types.Video, error) {
	videoList, err := appCtx.VideoModel.FindByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return FetchField(ctx, appCtx, videoList, selfId)
}

// QueryMultiVideo 根据给定的视频id列表获取视频信息
func QueryMultiVideo(ctx context.Context, appCtx *service.AppContext, selfId int64, videoIds []int64) ([]types.Video, error) {
	videoList, err := appCtx.VideoModel.FindMultiVideo(ctx, videoIds)
	if err != nil {
		return nil, err
	}
	return FetchField(ctx, appCtx, videoList, selfId)
}

func QueryByTimestamp(ctx context.Context, appCtx *service.AppContext, selfId int64, timestamp int64) (int64, []types.Video, error) {
	videoList, err := appCtx.VideoModel.FindByTimestamp(ctx, timestamp)
	var latestTime int64 = math.MaxInt64
	for _, v := range videoList {
		unixMicro := v.CreatedAt.UnixMilli()
		if unixMicro < latestTime {
			latestTime = unixMicro
		}
	}

	if err != nil {
		return 0, nil, err
	}
	videos, err := FetchField(ctx, appCtx, videoList, selfId)
	return latestTime, videos, err
}

func toVideo(video *dal.Videos) *types.Video {
	return &types.Video{
		CommentCount:  video.CommentCount,
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		ID:            video.Id,
		Title:         video.Title,
	}
}

func toVideos(videos []*dal.Videos) []types.Video {
	return itertool.Reduce(videos,
		func(agg []types.Video, item *dal.Videos, _ int) []types.Video {
			return append(agg, *toVideo(item))
		}, []types.Video{})
}

func videoToUids(videos []*dal.Videos) []int64 {
	return itertool.Reduce(videos, func(agg []int64, v *dal.Videos, i2 int) []int64 {
		return append(agg, v.UserId)
	}, nil)
}

// vIds 获取视频id列表
func vIds(videos []*dal.Videos) []int64 {
	return itertool.Reduce(videos, func(agg []int64, it *dal.Videos, _ int) []int64 {
		return append(agg, it.Id)
	}, make([]int64, 0, len(videos)))
}
