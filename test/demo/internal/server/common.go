package server

import (
	"context"
	pb "demo/api"
	"errors"
	"os"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

func DiscoveryService() (pb.DiscoveryClient, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dsAddr := os.Getenv("REGISTER_CENTER")
	if dsAddr == "" {
		var dc struct {
			Discovery *struct {
				Addrs string
			}
		}
		if err := paladin.Get("discovery.toml").UnmarshalTOML(&dc); err != nil {
			panic(err)
		}
		dsAddr = dc.Discovery.Addrs
	}
	log.Info("Register center address: %s", dsAddr)

	cli := warden.NewClient(nil)
	conn, err := cli.Dial(ctx, "direct://default/"+dsAddr)
	if err != nil {
		return nil, errors.New("Register center is unready")
	}

	return pb.NewDiscoveryClient(conn), nil
}
