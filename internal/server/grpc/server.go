package grpc

import (
	"register-center/internal/service"

	"context"
	pb "register-center/api"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server    *warden.ServerConfig
		Discovery *struct {
			Apipath string
		}
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	registerService(ws, svc, rc.Server.Addr, rc.Discovery.Apipath)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

func registerService(ws *warden.Server, svc *service.Service, addr string, apipath string) {
	pb.RegisterDiscoveryServer(ws.Server(), svc)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if _, err := svc.Register(ctx, &pb.RegSvcReqs{
		AppID: "discovery.service",
		Urls:  []string{addr},
	}); err != nil {
		panic(err)
	}
}
