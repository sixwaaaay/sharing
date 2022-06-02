package types

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

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
