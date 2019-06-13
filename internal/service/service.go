package service

import (
	"context"
	"math/rand"
	"os"
	pb "register-center/api"
	"register-center/internal/dao"
	"strings"
	"time"

	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/bilibili/kratos/pkg/log"
	"encoding/json"
)

type MicoService struct {
	Cancel context.CancelFunc
}

// Service service.
type Service struct {
	ac  *paladin.Map
	kong *dao.Kong
	sm  map[string]MicoService
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		kong: dao.NewKong(),
		sm:  make(map[string]MicoService),
	}
	return s
}

type DiscoveryConfig struct {
	Nodes   string
	Timeout xtime.Duration
}

func (s *Service) Register(ctx context.Context, req *pb.RegSvcReqs) (*pb.RegSvcResp, error) {
	// 将服务注册进Discovery中
	var dc struct {
		Discovery *DiscoveryConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&dc); err != nil {
		dc.Discovery.Nodes = "127.0.0.1:7171"
	}
	hn, _ := os.Hostname()
	dis := discovery.New(&discovery.Config{
		Nodes: strings.Split(dc.Discovery.Nodes, ","), // NOTE: 配置种子节点(1个或多个)，client内部可根据/discovery/nodes节点获取全部node(方便后面增减节点)
		Zone:  env.Zone,
		Env:   env.DeployEnv,
	})
	addrs := make([]string, len(req.Urls))
	for i, url := range req.Urls {
		addrs[i] = "grpc://" + strings.TrimLeft(url, "^(http|https)://")
	}
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    req.AppID,
		Hostname: hn,
		Addrs:    addrs,
	}
	if cancel, err := dis.Register(ctx, ins); err != nil {
		return nil, err
	} else {
		log.Info("Service %s has registered into discovery", req.AppID)
		s.sm[req.AppID] = MicoService{
			cancel,
		}
	}
	// 将服务注册进Kong中
	if svcID, err := s.kong.NewService(req.AppID, req.Urls); err != nil {
		return nil, err
	} else {
		return &pb.RegSvcResp{ KongID: svcID }, nil
	}
}

func (s *Service) Services(context.Context, *empty.Empty) (resp *pb.LstSvcResp, err error) {
	for appID := range s.sm {
		resp.AppIDs = append(resp.AppIDs, appID)
	}
	return
}

type consumer struct {
	conf    *discovery.Config
	timeout xtime.Duration
	appID   string
	dis     naming.Resolver
	ins     []*naming.Instance
}

func (csm *consumer) getInstances(ctx context.Context, ch <-chan struct{}) {
	for _, ok := <-ch; ok; _, ok = <-ch {
		if ins, ok := csm.dis.Fetch(ctx); !ok {
			continue
		} else if in, ok := ins.Instances[csm.conf.Zone]; ok {
			csm.ins = in
		} else {
			for _, in := range ins.Instances {
				csm.ins = append(csm.ins, in...)
			}
		}
	}
}

func (csm *consumer) getInstance() *naming.Instance {
	logTime := time.Now()
	dur := time.Duration(csm.timeout)
	for len(csm.ins) == 0 {
		if time.Since(logTime) > dur {
			return nil
		} else {
			time.Sleep(2 * time.Second)
		}
	}
	// NOTE: 此处运用一种负载均衡算法得出一个实例用于处理
	rand.Seed(time.Now().Unix())
	return csm.ins[rand.Intn(len(csm.ins))]
}

func (s *Service) Service(ctx context.Context, req *pb.IdenSvcReqs) (*pb.GetSvcResp, error) {
	var dc struct {
		Discovery *DiscoveryConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&dc); err != nil {
		dc.Discovery.Nodes = "127.0.0.1:7171"
		dc.Discovery.Timeout = xtime.Duration(5 * time.Second)
	}
	cfg := &discovery.Config{
		Nodes: strings.Split(dc.Discovery.Nodes, ","),
		Zone:  env.Zone,
		Env:   env.DeployEnv,
	}
	dis := discovery.New(cfg)
	csm := &consumer{
		cfg,
		dc.Discovery.Timeout,
		req.AppID,
		dis.Build(req.AppID),
		nil,
	}
	ch := csm.dis.Watch()
	go csm.getInstances(ctx, ch)
	ins := csm.getInstance()
	return &pb.GetSvcResp{Addrs: ins.Addrs}, nil
}

func (s *Service) Cancel(ctx context.Context, req *pb.IdenSvcReqs) (*empty.Empty, error) {
	if svc, exs := s.sm[req.AppID]; !exs {
		return nil, errors.New("未找到指定服务，取消失败")
	} else {
		svc.Cancel()
	}
	return &empty.Empty{}, nil
}

func (s *Service) AddRoutes(ctx context.Context, req *pb.AddRoutesReqs) (resp *pb.AddRoutesResp, err error) {
	paths := make(map[string]interface{})
	if err := json.Unmarshal(req.Paths, &paths); err != nil {
		return nil, err
	}
	resp = &pb.AddRoutesResp{}
	for path, body := range paths {
		for method, inbody := range body.(map[string]interface{}) {
			// NOTE: 默认用summary的最后一截作为路由的名字，所以不能包含特殊字符
			summary := inbody.(map[string]interface{})["summary"].(string)
			nameArray := strings.Split(summary, "/")
			resp.Routes = append(resp.Routes, &pb.RouteResp{
				Name: nameArray[len(nameArray) - 1],
				Path: path,
				Method: method,
			})
		}
	}
	for i, route := range resp.Routes {
		if rid, err := s.kong.AddRoute(route.Name, route.Method, route.Path); err != nil {
			return nil, err
		} else {
			resp.Routes[i].Id = rid
		}
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.kong.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	for appID, svc := range s.sm {
		svc.Cancel()
		log.Info("Service %s canceled from discovery", appID)
	}
}
