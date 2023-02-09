package types

import (
	"mime/multipart"
)

type UserReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type UserInfoReq struct {
	UserID int64  `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type UserInfoResp struct {
	StatusCode int64   `json:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`     // 返回状态描述
	User       *User   `json:"user,omitempty"` // 用户信息
}

type FeedReq struct {
	LatestTime *int64  `form:"latest_time" json:"latest_time,omitempty"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `form:"token" json:"token,omitempty"`             // 登录状态则有
}

type FeedResp struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

type UploadReq struct {
	Token string                `form:"token" json:"token" binding:"required"`
	Title string                `form:"title" json:"title" binding:"required"`
	File  *multipart.FileHeader `form:"data" binding:"required"`
} // 经测试可以完成绑定

type UploadResp struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
}

type PubListReq struct {
	Token  string `form:"token" json:"token"`
	UserId int64  `form:"user_id" json:"user_id"`
}

type PubListResp struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

type FollowActionReq struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int64  `form:"action_type" json:"action_type" binding:"required"`
}

type FollowActionResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type RelationReq struct {
	UserId int64  `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type FollowListResp struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户信息列表
}

type FollowerListResp struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户信息列表
}

type FavoriteReq struct {
	UserId     int64  `form:"user_id"`     // 用户id
	Token      string `form:"token"`       // 用户token
	VideoId    int64  `form:"video_id"`    // 视频id
	ActionType int64  `form:"action_type"` //1-点赞，2-取消点赞
}

type FavoriteResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteListReq struct {
	UserId int64  `form:"user_id"` // 用户id
	Token  string `form:"token"`   // 用户token
}

type FavoriteListResp struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}

type CommentReq struct {
	UserId      int64  `form:"user_id" json:"user_id"`           // 用户id
	Token       string `form:"token" json:"token"`               // 用户token
	VideoId     int64  `form:"video_id" json:"video_id"`         // 视频id
	ActionType  int8   `form:"action_type" json:"action_type"`   // 1- 发布评论 2- 删除
	CommentText string `form:"comment_text" json:"comment_text"` // 评论内容
	CommentID   int64  `form:"comment_id" json:"comment_id"`     // 评论id
}

type CommentResp struct {
	Comment    *Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg"`  // 返回状态描述
}

type CommentListReq struct {
	Token   string `form:"token" json:"token"`       // 用户token
	VideoId int64  `form:"video_id" json:"video_id"` // 视频id
}

type CommentListResp struct {
	CommentList []Comment `json:"comment_list"` // 评论列表
	StatusCode  int64     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

// Video 视频信息
type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// Comment 评论
type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

// User 用户信息
type User struct {
	FollowCount   int64  `json:"follow_count"`     // 关注总数
	FollowerCount int64  `json:"follower_count"`   // 粉丝总数
	ID            int64  `json:"id"`               // 用户id
	IsFollow      bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name          string `json:"name"`             // 用户名称
	Avatar        string `json:"avatar,omitempty"` // 用户头像 url
}
