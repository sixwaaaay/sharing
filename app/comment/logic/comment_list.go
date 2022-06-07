package logic

import (
	"bytelite/app/basic"
	"bytelite/app/comment/types"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
)

type CommentListLogic func(req *types.CommentListReq) (*types.CommentListResp, error)

var NewCommentListLogic = newCommentListLogic

func newCommentListLogic(ctx context.Context, appCtx *service.AppContext) CommentListLogic {
	return func(req *types.CommentListReq) (*types.CommentListResp, error) {
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		comments, err := basic.QueryVideoComment(ctx, appCtx, selfId, req.VideoId)
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
