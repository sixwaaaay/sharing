package logic

import (
	"bytelite/app/basic"
	"bytelite/app/feed/types"
	"bytelite/common/errorx"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FeedLogic func(req *types.FeedReq) (*types.FeedResp, error)

var NewFeedLogic = newFeedLogic

func newFeedLogic(ctx context.Context, appCtx *service.AppContext) FeedLogic {
	logger := logx.WithContext(ctx)
	return func(req *types.FeedReq) (*types.FeedResp, error) {
		var timestamp int64
		if req.LatestTime == nil {
			timestamp = time.Now().UnixMilli()
		} else {
			timestamp = *req.LatestTime
		}
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		unixStamp, videoList, err := basic.QueryByTimestamp(ctx, appCtx, selfId, timestamp)

		if err != nil {
			logger.Errorf("query video by timestamp failed, err: %+v", err)
			return nil, errorx.NewDefaultError("query video list failed")
		}
		return &types.FeedResp{
			NextTime:  &unixStamp,
			VideoList: videoList,
		}, nil
	}
}
