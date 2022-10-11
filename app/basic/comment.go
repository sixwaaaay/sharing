package basic

import (
	"bytelite/app/comment/dal"
	"bytelite/common/cotypes"
	"bytelite/common/covert"
	"bytelite/common/errorx"
	"bytelite/common/itertool"
	"bytelite/service"
	"context"
)

// RemoveComment 移除评论
func RemoveComment(ctx context.Context, appCtx *service.AppContext, selfId, commentId, _ int64) error {
	err := appCtx.CommentsModel.DeleteUserComment(ctx, selfId, commentId)
	if err != nil {
		return err
	}
	return err
}

// AddComment 添加评论
func AddComment(ctx context.Context, appCtx *service.AppContext, userId int64, videoId int64, content string) (*cotypes.Comment, error) {
	comment := &dal.Comments{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}
	ret, err := appCtx.CommentsModel.Insert(ctx, comment)
	if err != nil {
		return nil, errorx.NewDefaultError("comment failed")
	}
	comment.Id, err = ret.LastInsertId()
	if err != nil {
		return nil, errorx.NewDefaultError("server busy")
	}
	return toComment(comment), nil
}

// QueryVideoComment 查询视频的评论
func QueryVideoComment(ctx context.Context, appCtx *service.AppContext, selfId, videoId int64) ([]cotypes.Comment, error) {
	commentsList, err := appCtx.CommentsModel.FindCommentsByVideoID(ctx, videoId)
	if err != nil {
		return nil, errorx.NewDefaultError("query failed, server busy")
	}
	// 查询用户信息
	multiUserInfo, err := QueryMultiUserInfo(ctx, appCtx, selfId, commentsToUserIds(commentsList))
	userMap := UserMap(multiUserInfo)
	comments := toComments(commentsList)
	// 填充评论的用户信息
	for i := 0; i < len(comments); i++ {
		if user, ok := userMap[commentsList[i].UserId]; ok {
			comments[i].User = user
		}
	}
	return comments, nil
}

func toComment(comment *dal.Comments) *cotypes.Comment {
	return &cotypes.Comment{
		Content:    comment.Content,
		CreateDate: covert.TimeFormatMMDD(comment.CreatedAt),
		ID:         comment.Id,
	}
}

func toComments(comments []*dal.Comments) []cotypes.Comment {
	return itertool.Reduce(comments, func(agg []cotypes.Comment, item *dal.Comments, _ int) []cotypes.Comment {
		return append(agg, *toComment(item))
	}, []cotypes.Comment{})
}

func commentsToUserIds(comments []*dal.Comments) []int64 {
	return itertool.Reduce(comments, func(agg []int64, item *dal.Comments, _ int) []int64 {
		return append(agg, item.UserId)
	}, []int64{})
}

func UserMap(users []cotypes.User) map[int64]cotypes.User {
	userMap := make(map[int64]cotypes.User, len(users))
	for _, u := range users {
		userMap[u.ID] = u
	}
	return userMap
}
