package etc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"strconv"
)

type Config struct {
	rest.RestConf
	DSN       string
	JWTSecret string
	Cache     cache.CacheConf
	Minio     struct {
		Endpoint  string `json:",default=http://localhost:9000"`
		AccessKey string `json:",default=minio"`
		SecretKey string `json:",default=minio123"`
		UseSSL    bool   `json:",default=false"`
		Bucket    string `json:",default=zero"`
	}
	ContentBaseUrl string
}

func (c *Config) Addr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
