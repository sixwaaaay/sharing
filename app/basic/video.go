package basic

import (
	"bytelite/app/publish/dal"
	"bytelite/common/cotypes"
	"bytelite/common/covert"
	"bytelite/service"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

// FetchField 填充相关字段
func FetchField(ctx context.Context, appCtx *service.AppContext, videoList []*dal.Videos, selfId int64) ([]cotypes.Video, error) {
	logger := logx.WithContext(ctx)
	ids := videoToUids(videoList)
	multiUserInfo, err := QueryMultiUserInfo(ctx, appCtx, selfId, ids)
	if err != nil {
		logger.Errorf("QueryMultiUserInfo error: %v", err)
		return nil, err
	}
	userMap := covert.UserMap(multiUserInfo)
	logx.Infof("userMap: %v", userMap)
	videos := toVideos(videoList)
	favorites, err := QueryVideoFavorites(ctx, appCtx, selfId, vIds(videoList))
	if err != nil {
		logger.Errorf("QueryVideoFavorites error: %v", err)
		return nil, err
	}

	favoriteMap := covert.Int64SliceToMap(favorites)
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
func QueryUserVideo(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) ([]cotypes.Video, error) {
	videoList, err := appCtx.VideoModel.FindByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return FetchField(ctx, appCtx, videoList, selfId)
}

// QueryMultiVideo 根据给定的视频id列表获取视频信息
func QueryMultiVideo(ctx context.Context, appCtx *service.AppContext, selfId int64, videoIds []int64) ([]cotypes.Video, error) {
	videoList, err := appCtx.VideoModel.FindMultiVideo(ctx, videoIds)
	if err != nil {
		return nil, err
	}
	return FetchField(ctx, appCtx, videoList, selfId)
}

func QueryByTimestamp(ctx context.Context, appCtx *service.AppContext, selfId int64, timestamp int64) (int64, []cotypes.Video, error) {
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

func UpdateVideoCommentCount(ctx context.Context, appCtx *service.AppContext, videoId int64, count int64) error {
	_, err := appCtx.VideoModel.UpdateCommentCount(ctx, videoId, count)
	return err
}

func UpdateVideoFavoriteCount(ctx context.Context, appCtx *service.AppContext, videoId int64, count int64) error {
	_, err := appCtx.VideoModel.UpdateFavoriteCount(ctx, videoId, count)
	return err
}

func toVideo(video *dal.Videos) *cotypes.Video {
	return &cotypes.Video{
		CommentCount:  video.CommentCount,
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		ID:            video.Id,
		Title:         video.Title,
	}
}

func toVideos(videos []*dal.Videos) []cotypes.Video {
	var videoList []cotypes.Video
	for _, video := range videos {
		videoList = append(videoList, *toVideo(video))
	}
	return videoList
}

func videoToUids(videos []*dal.Videos) []int64 {
	ids := make([]int64, 0, len(videos))
	for _, c := range videos {
		ids = append(ids, c.UserId)
	}
	return ids
}

// vIds 获取视频id列表
func vIds(videos []*dal.Videos) []int64 {
	ids := make([]int64, 0, len(videos))
	for _, c := range videos {
		ids = append(ids, c.Id)
	}
	return ids
}
