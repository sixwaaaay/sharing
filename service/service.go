package service

import (
	comments "bytelite/app/comment/dal"
	favorites "bytelite/app/favorite/dal"
	videos "bytelite/app/publish/dal"
	relation "bytelite/app/relation/dal"
	users "bytelite/app/user/dal"
	"bytelite/common/auth"
	"bytelite/etc"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AppContext struct {
	RelationModel  relation.RelationsModel
	FavoriteModel  favorites.FavoritesModel
	UsersModel     users.UsersModel
	CommentsModel  comments.CommentsModel
	VideoModel     videos.VideosModel
	JWTSigner      *auth.JWTSigner
	MinioClient    *minio.Client
	MinioBucket    string
	ContentBaseUrl string
}

func NewAppContext(c *etc.Config) *AppContext {
	sqlConn := sqlx.NewMysql(c.DSN)
	client, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		panic(err)
	}
	return &AppContext{
		RelationModel:  relation.NewRelationsModel(sqlConn, c.Cache),
		FavoriteModel:  favorites.NewFavoritesModel(sqlConn, c.Cache),
		UsersModel:     users.NewUserModel(sqlConn, c.Cache),
		CommentsModel:  comments.NewCommentsModel(sqlConn, c.Cache),
		VideoModel:     videos.NewVideosModel(sqlConn, c.Cache),
		JWTSigner:      auth.NewJWTSigner(c.JWTSecret),
		MinioClient:    client,
		MinioBucket:    c.Minio.Bucket,
		ContentBaseUrl: c.ContentBaseUrl,
	}
}

func (c *AppContext) DefaultCoverUrl() string {
	return fmt.Sprintf("/%s/default.png", c.MinioBucket)
}
