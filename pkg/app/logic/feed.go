package logic

import (
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"context"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/common/middleware"
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
		unixStamp, videoList, err := QueryByTimestamp(ctx, appCtx, selfId, timestamp)

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
