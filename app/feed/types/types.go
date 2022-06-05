package types

import "bytelite/common/cotypes"

type FeedReq struct {
	LatestTime *int64  `form:"latest_time" json:"latest_time,omitempty"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `form:"token" json:"token,omitempty"`             // 登录状态则有
}

type FeedResp struct {
	NextTime   *int64          `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `json:"status_msg"`  // 返回状态描述
	VideoList  []cotypes.Video `json:"video_list"`  // 视频列表
}
