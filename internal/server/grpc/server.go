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
	pb.RegisterRegisterServer(ws.Server(), svc)
	if _, err := svc.RegAsGRPC(context.Background(), &pb.RegSvcReqs{
		AppID: "register.service",
		Urls:  []string{"127.0.0.1:9093"}, // 这里给出的url很关键，必须是从外部可以访问的
	}); err != nil {
		panic(err)
	}
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
