package logic

import (
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"

	"context"
	"github.com/sixwaaaay/sharing/common/middleware"
)

type CommentListLogic func(req *types.CommentListReq) (*types.CommentListResp, error)

var NewCommentListLogic = newCommentListLogic

func newCommentListLogic(ctx context.Context, appCtx *service.AppContext) CommentListLogic {
	return func(req *types.CommentListReq) (*types.CommentListResp, error) {
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		comments, err := QueryVideoComment(ctx, appCtx, selfId, req.VideoId)
		if err != nil {
			return nil, err
		}
		return &types.CommentListResp{
			CommentList: comments,
			StatusCode:  0,
			StatusMsg:   nil,
		}, nil
	}
}
