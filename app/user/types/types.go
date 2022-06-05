package types

import "bytelite/common/cotypes"

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
	StatusCode int64         `json:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMsg  *string       `json:"status_msg"`     // 返回状态描述
	User       *cotypes.User `json:"user,omitempty"` // 用户信息
}
