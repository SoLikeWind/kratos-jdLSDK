package data

import (
	"helloworld/internal/conf"
	"helloworld/internal/pkg/jdl"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewJdLConfig)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewJdLConfig(conf *conf.JdL, logger log.Logger) *jdl.Config {
	return jdl.NewJdConfig(
		jdl.WithAppKey(conf.AppKey),
		jdl.WithAppSecret(conf.AppSecret),
		jdl.WithAccessToken(conf.AccessToken),
		jdl.WithEnv(conf.Env),
		jdl.WithLopDn(conf.LOP_DN),
	)
}
