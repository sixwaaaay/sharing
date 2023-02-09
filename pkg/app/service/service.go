package service

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sixwaaaay/sharing/configs"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/common/auth"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AppContext struct {
	RelationModel  dal.RelationsModel
	FavoriteModel  dal.FavoritesModel
	UsersModel     dal.UsersModel
	CommentsModel  dal.CommentsModel
	VideoModel     dal.VideosModel
	JWTSigner      *auth.JWTSigner
	MinioClient    *minio.Client
	MinioBucket    string
	ContentBaseUrl string
}

func NewAppContext(c *configs.Config) *AppContext {
	sqlConn := sqlx.NewMysql(c.DSN)
	client, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		panic(err)
	}
	return &AppContext{
		RelationModel:  dal.NewRelationsModel(sqlConn, c.Cache),
		FavoriteModel:  dal.NewFavoritesModel(sqlConn, c.Cache),
		UsersModel:     dal.NewUserModel(sqlConn, c.Cache),
		CommentsModel:  dal.NewCommentsModel(sqlConn, c.Cache),
		VideoModel:     dal.NewVideosModel(sqlConn, c.Cache),
		JWTSigner:      auth.NewJWTSigner(c.JWTSecret),
		MinioClient:    client,
		MinioBucket:    c.Minio.Bucket,
		ContentBaseUrl: c.ContentBaseUrl,
	}
}

func (c *AppContext) DefaultCoverUrl() string {
	return fmt.Sprintf("/%s/default.png", c.MinioBucket)
}
