package types

import "bytelite/common/cotypes"

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
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `json:"status_msg"`  // 返回状态描述
	VideoList  []cotypes.Video `json:"video_list"`  // 用户点赞视频列表
}
