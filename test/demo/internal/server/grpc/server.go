package grpc

import (
	"context"
	pb "demo/api"
	"demo/internal/server"
	"demo/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	registerService(ws, svc, rc.Server.Addr)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

func registerService(ws *warden.Server, svc *service.Service, addr string) {
	pb.RegisterDemoServer(ws.Server(), svc)

	if dsSvc, err := server.DiscoveryService(); err != nil {
		log.Error("Fetch discovery service error: %v", err)
	} else if _, err := dsSvc.Register(context.Background(), &pb.RegSvcReqs{
		AppID: "demo.service",
		Urls:  []string{addr},
	}); err != nil {
		panic(err)
	}
}
