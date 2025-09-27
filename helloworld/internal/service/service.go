package service

import (
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"helloworld/internal/pkg/jdl"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewCreateJdLClient)

func NewCreateJdLClient(conf *conf.JdL, logger log.Logger) (*jdl.JdLClient, func(), error) {
	createJdLClient, _, err := jdl.NewJdClient(jdl.NewJdConfig(
		jdl.WithAppKey(conf.AppKey),
		jdl.WithAppSecret(conf.AppSecret),
		jdl.WithAccessToken(conf.AccessToken),
		jdl.WithEnv(conf.Env),
		jdl.WithLopDn(conf.LOP_DN),
	), logger)
	if err != nil {
		return nil, nil, err
	}

	return createJdLClient, func() {}, nil
}

// type Service struct {
// 	JdLClient *jdl.JdLClient
// }

// func NewService(jdLClient *jdl.JdLClient) *Service {
// 	return &Service{JdLClient: jdLClient}
// }
