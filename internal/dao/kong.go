package dao

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"register-center/internal/utils"
	"strings"
	"time"

	"encoding/json"
	"io/ioutil"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

type Kong struct {
	adminAddrs []string
}

func NewKong() *Kong {
	var kc struct {
		Base struct {
			AdminAddrs string
		}
	}
	if adminAddrs := os.Getenv("KONG_ADMIN_ADDR"); len(adminAddrs) != 0 {
		kc.Base.AdminAddrs = adminAddrs
	} else if err := paladin.Get("kong.toml").UnmarshalTOML(&kc); err != nil {
		panic(err)
	}
	return &Kong{
		adminAddrs: strings.Split(kc.Base.AdminAddrs, ","),
	}
}

func (k *Kong) GetAdminAddr(ps ...string) string {
	if len(k.adminAddrs) == 0 {
		return ""
	}
	rand.Seed(time.Now().Unix())
	p := ""
	if len(ps) != 0 {
		p = path.Join(ps...)
	}
	return k.adminAddrs[rand.Intn(len(k.adminAddrs))] + p
}

// 一个服务（service）对应一个上游（upstream）；一个上游（upstream）对应多个转发（target）
func (k *Kong) NewService(name string, urls []string) (string, error) {
	// 创建上游转发组
	data := []byte(fmt.Sprintf("{\"name\":\"%s\"}", name))
	if _, err := http.Post(k.GetAdminAddr("/upstreams"), utils.APPLICATION_JSON, bytes.NewReader(data)); err != nil {
		return "", err
	} else {
		log.Info("Upstream %s has been created", name)
	}
	upTgtUrl := fmt.Sprintf("/upstreams/%s/targets", name)
	for _, url := range urls {
		// 采用默认权重，都是100
		data = []byte(fmt.Sprintf("{\"target\":\"%s\"}", url))
		if _, err := http.Post(k.GetAdminAddr(upTgtUrl), utils.APPLICATION_JSON, bytes.NewReader(data)); err != nil {
			return "", err
		} else {
			log.Info("Upstream target %s has been created", url)
		}
	}

	// 新建服务
	data = []byte(fmt.Sprintf("{\"name\":\"%s\", \"host\":\"%s\"}", name, name))
	if resp, err := http.Post(k.GetAdminAddr("/services"), utils.APPLICATION_JSON, bytes.NewReader(data)); err != nil {
		return "", err
	} else if byteResp, err := ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()

		result := make(map[string]interface{})
		if err := json.Unmarshal(byteResp, &result); err != nil {
			return "", err
		} else if id, exs := result["id"]; !exs {
			return "", fmt.Errorf("返回体中没有服务ID，是否返回体结构有变：%s", string(byteResp))
		} else {
			log.Info("Service %s has been created, id is %s", name, id)
			// 路由应该只开http路由
			return id.(string), nil
		}
	}
}

func (k *Kong) AddRoute(svc string, name string, method string, path string) (string, error) {
	data := []byte(fmt.Sprintf("{\"name\":\"%s\", \"methods\":[\"%s\"], \"paths\":[\"%s\"]}", name, strings.ToUpper(method), path))
	if resp, err := http.Post(k.GetAdminAddr(fmt.Sprintf("/services/%s/routes", svc)), utils.APPLICATION_JSON, bytes.NewReader(data)); err != nil {
		return "", err
	} else if byteResp, err := ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()

		result := make(map[string]interface{})
		if err := json.Unmarshal(byteResp, &result); err != nil {
			return "", err
		} else if id, exs := result["id"]; !exs {
			return "", fmt.Errorf("返回体中没有服务ID，是否返回体结构有变：%s", string(byteResp))
		} else {
			log.Info("Route %s has been created, id is %s", name, id)
			return id.(string), nil
		}
	}
}

func (k *Kong) Ping(ctx context.Context) error {
	for _, addr := range k.adminAddrs {
		if _, err := http.Get(addr); err != nil {
			return fmt.Errorf("管理节点：%s连接异常：%v", addr, err)
		}
	}
	return nil
}
