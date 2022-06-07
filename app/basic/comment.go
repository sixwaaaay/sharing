package basic

import (
	"bytelite/app/comment/dal"
	"bytelite/common/cotypes"
	"bytelite/common/covert"
	"bytelite/service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// RemoveComment 移除评论
func RemoveComment(ctx context.Context, appCtx *service.AppContext, selfId, commentId, videoId int64) error {
	err := appCtx.CommentsModel.DeleteUserComment(ctx, selfId, commentId)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return
		}
		err = UpdateVideoCommentCount(ctx, appCtx, videoId, -1)
		if err != nil {
			logx.WithContext(ctx).Errorf("update video comment count failed, videoId: %d, err: %v", 0, err)
		}
	}()
	return err
}

// AddComment 添加评论
func AddComment(ctx context.Context, appCtx *service.AppContext, userId int64, videoId int64, content string) (*cotypes.Comment, error) {
	comment := &dal.Comments{
		UserId:    userId,
		VideoId:   videoId,
		Content:   content,
		CreatedAt: time.Now(),
	}
	ret, err := appCtx.CommentsModel.Insert(ctx, comment)
	if err != nil {
		return nil, err
	}
	// 添加视频的评论数量,可以失败，因此使用defer
	defer func() {
		if err != nil {
			return
		}
		err = UpdateVideoCommentCount(ctx, appCtx, videoId, 1)
		if err != nil {
			logx.WithContext(ctx).Errorf("update video comment count failed, videoId: %d, err: %v", videoId, err)
		}
	}()
	commentId, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}

	comment.Id = commentId

	return toComment(comment), nil
}

func QueryVideoComment(ctx context.Context, appCtx *service.AppContext, selfId, videoId int64) ([]cotypes.Comment, error) {
	commentsList, err := appCtx.CommentsModel.FindCommentsByVideoID(ctx, videoId)
	if err != nil {
		return nil, err
	}
	// 查询用户信息
	multiUserInfo, err := QueryMultiUserInfo(ctx, appCtx, selfId, commentsToUids(commentsList))
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
	ret := make([]cotypes.Comment, 0, len(comments))
	for _, comment := range comments {
		ret = append(ret, *toComment(comment))
	}
	return ret
}

func commentsToUids(comments []*dal.Comments) []int64 {
	ids := make([]int64, 0, len(comments))
	for _, c := range comments {
		ids = append(ids, c.UserId)
	}
	return ids
}

func UserMap(users []cotypes.User) map[int64]cotypes.User {
	userMap := make(map[int64]cotypes.User, len(users))
	for _, u := range users {
		userMap[u.ID] = u
	}
	return userMap
}
