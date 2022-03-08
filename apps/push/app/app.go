package app

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"github.com/yusank/goim/pkg/db/redis"
	"github.com/yusank/goim/pkg/registry"
)

type Application struct {
	Core           *kratos.App
	Register       registry.RegisterDiscover
	ServerConfig   *Config
	RegisterConfig *Registry
	Redis          *redisv8.Client
	agentID        string
}

var (
	application *Application
	onceChan    = make(chan struct{}, 1)
)

func InitApplication(confPath string) (*Application, error) {
	// only can call this func once, if call twice will be panic
	onceChan <- struct{}{}
	close(onceChan)

	cfg, regCfg := ParseConfig(confPath)
	//s := &service.PushMessager{}

	var servers = make([]transport.Server, 0)
	if cfg.Http != nil {
		httpSrv := http.NewServer(
			http.Address(fmt.Sprintf("%s:%d", cfg.Http.GetAddr(), cfg.Http.GetPort())),
			http.Middleware(
				recovery.Recovery(),
			),
		)
		servers = append(servers, httpSrv)
	}
	if cfg.Grpc != nil {
		// services
		grpcSrv := grpc.NewServer(
			grpc.Address(fmt.Sprintf("%s:%d", cfg.Grpc.GetAddr(), cfg.Grpc.GetPort())),
			grpc.Middleware(
				recovery.Recovery(),
			),
		)
		servers = append(servers, grpcSrv)
	}

	var options = []kratos.Option{
		kratos.Name(cfg.GetName()),
		kratos.Version(cfg.GetVersion()),
		kratos.Server(
			servers...,
		),
		kratos.Metadata(cfg.GetMetadata()),
	}

	reg, err := registry.NewRegistry(regCfg.Registry)
	if err != nil {
		return nil, err
	}
	if reg != nil {
		options = append(options, kratos.Registrar(reg))
	}

	rdb, err := redis.NewRedis(redis.WithConfig(cfg.GetRedis()))
	if err != nil {
		return nil, err
	}

	core := kratos.New(
		options...,
	)

	application = &Application{
		Core:           core,
		ServerConfig:   cfg,
		RegisterConfig: regCfg,
		Register:       reg,
		Redis:          rdb,
		agentID:        uuid.NewString(),
	}

	return application, nil
}

func (a *Application) Run() error {
	return a.Core.Run()
}

func GetApplication() *Application {
	return application
}

func GetAgentID() string {
	return application.agentID
}
