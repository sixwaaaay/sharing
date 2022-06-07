package logic

import (
	"bytelite/app/publish/dal"
	"bytelite/app/publish/types"
	"bytelite/common/errorx"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
)

type UploadLogic func(req *types.UploadReq) (*types.UploadResp, error)

var NewUploadLogic = newUploadLogic

func newUploadLogic(ctx context.Context, appCtx *service.AppContext) UploadLogic {
	return func(req *types.UploadReq) (*types.UploadResp, error) {
		logger := logx.WithContext(ctx)
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		open, err := req.File.Open()
		if err != nil {
			return nil, errorx.NewDefaultError("open file error")
		}
		defer open.Close()
		// 生成文件名, 是否要加上文件后缀?
		suffix, err := getFileType(req.File.Filename)
		if err != nil {
			logger.Errorf("get file type error: %v", err)
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
				UserId:   selfId,
				PlayUrl:  fmt.Sprintf("/%s/%s", appCtx.MinioBucket, name),
				CoverUrl: appCtx.DefaultCoverUrl(), // default
				Title:    req.Title,
				// todo: 时间要设置自动启用默认值
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
