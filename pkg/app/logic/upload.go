package logic

import (
	"github.com/minio/minio-go/v7"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/common/middleware"

	"context"
	"fmt"
	"strings"
	"time"
)

type UploadLogic func(req *types.UploadReq) (*types.UploadResp, error)

var NewUploadLogic = newUploadLogic

func newUploadLogic(ctx context.Context, appCtx *service.AppContext) UploadLogic {
	return func(req *types.UploadReq) (*types.UploadResp, error) {
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		open, err := req.File.Open()
		if err != nil {
			return nil, errorx.NewDefaultError("open file error")
		}
		defer open.Close()
		suffix, err := getFileType(req.File.Filename)
		if err != nil {
			return nil, errorx.NewDefaultError("get file type error")
		}
		name := fmt.Sprintf("%d%d.%s", selfId, time.Now().UnixMilli(), suffix)
		// 写入对象储存
		_, err = appCtx.MinioClient.PutObject(ctx, appCtx.MinioBucket, name, open, req.File.Size, minio.PutObjectOptions{})
		if err != nil {
			return nil, errorx.NewDefaultError("upload file error")
		}
		// 写入数据库
		_, err = appCtx.VideoModel.Insert(ctx,
			&dal.Videos{
				UserId:    selfId,
				PlayUrl:   fmt.Sprintf("/%s/%s", appCtx.MinioBucket, name),
				CoverUrl:  appCtx.DefaultCoverUrl(), // default
				Title:     req.Title,
				CreatedAt: time.Now(),
			},
		)
		if err != nil {
			return nil, errorx.NewDefaultError("insert video error")
		}

		return &types.UploadResp{
			StatusCode: 0,
		}, nil
	}
}

func getFileType(fileName string) (string, error) {
	// get the "." index
	index := strings.LastIndex(fileName, ".")
	if index == -1 {
		return "", errorx.NewDefaultError("不支持的文件类型")
	}
	return fileName[index+1:], nil
}
