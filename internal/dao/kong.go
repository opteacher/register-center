package dao

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"fmt"
	"context"
)

type Kong struct {

}

func NewKong() *Kong {
	var kc struct{
		Base struct{
			AdminAddrs string
		}
	}
	if err := paladin.Get("kong.toml").UnmarshalTOML(&kc); err != nil {
		panic(err)
	}
	fmt.Println(kc)
	return &Kong{}
}

func (k *Kong) Ping(ctx context.Context) error {
	return nil
}