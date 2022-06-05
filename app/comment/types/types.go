package types

import "bytelite/common/cotypes"

type CommentReq struct {
	UserId      int64  `form:"user_id" json:"user_id"`           // 用户id
	Token       string `form:"token" json:"token"`               // 用户token
	VideoId     int64  `form:"video_id" json:"video_id"`         // 视频id
	ActionType  int8   `form:"action_type" json:"action_type"`   // 1- 发布评论 2- 删除
	CommentText string `form:"comment_text" json:"comment_text"` // 评论内容
	CommentID   int64  `form:"comment_id" json:"comment_id"`     // 评论id
}

type CommentResp struct {
	Comment    *cotypes.Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string          `json:"status_msg"`  // 返回状态描述
}

type CommentListReq struct {
	Token   string `form:"token" json:"token"`       // 用户token
	VideoId int64  `form:"video_id" json:"video_id"` // 视频id
}

type CommentListResp struct {
	CommentList []cotypes.Comment `json:"comment_list"` // 评论列表
	StatusCode  int64             `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string           `json:"status_msg"`   // 返回状态描述
}
