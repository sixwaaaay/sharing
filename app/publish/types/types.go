package types

import (
	"bytelite/common/cotypes"
	"mime/multipart"
)

type UploadReq struct {
	Token string                `form:"token" json:"token"`
	Title string                `form:"title" json:"title"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
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
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `json:"status_msg"`  // 返回状态描述
	VideoList  []cotypes.Video `json:"video_list"`  // 用户发布的视频列表
}
