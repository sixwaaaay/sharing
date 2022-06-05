package types

import "bytelite/common/cotypes"

type FollowActionReq struct {
	UserId     int64  `form:"user_id" json:"user_id" binding:"required"`
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
	StatusCode int64          `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string        `json:"status_msg"`  // 返回状态描述
	UserList   []cotypes.User `json:"user_list"`   // 用户信息列表
}

type FollowerListResp struct {
	StatusCode int64          `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string        `json:"status_msg"`  // 返回状态描述
	UserList   []cotypes.User `json:"user_list"`   // 用户信息列表
}
