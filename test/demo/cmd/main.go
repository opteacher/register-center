package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo/internal/server/grpc"
	"demo/internal/server/http"
	"demo/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("demo start")
	svc := service.New()
	grpcSrv := grpc.New(svc)
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := grpcSrv.Shutdown(ctx); err != nil {
				log.Error("grpcSrv.Shutdown error(%v)", err)
			}
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("demo exit")
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
