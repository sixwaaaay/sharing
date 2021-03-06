package logic

import (
	"bytelite/app/basic"
	"bytelite/app/comment/types"
	"bytelite/common/errorx"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
)

type CommentLogic func(req *types.CommentReq) (*types.CommentResp, error)

var NewCommentLogic = newCommentLogic

func newCommentLogic(ctx context.Context, appCtx *service.AppContext) CommentLogic {
	return func(req *types.CommentReq) (*types.CommentResp, error) {
		resp := &types.CommentResp{}
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		switch req.ActionType {
		case 1:
			if req.CommentText == "" {
				return nil, errorx.NewDefaultError("评论内容不能为空")
			}
			comment, err := basic.AddComment(ctx, appCtx, selfId, req.VideoId, req.CommentText)
			if err != nil {
				return nil, errorx.NewDefaultError("add comment error")
			}
			resp.StatusCode = 0
			resp.Comment = comment
			return resp, nil
		case 2:
			err := basic.RemoveComment(ctx, appCtx, selfId, req.CommentID, req.VideoId)
			if err != nil {
				return nil, errorx.NewDefaultError("delete comment error")
			}
			resp.StatusCode = 0
			return resp, nil
		default:
			return nil, errorx.NewDefaultError("unsupported action type")
		}
	}
}
