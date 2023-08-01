package manager

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type (
	ComponentPlugin interface {
		// Name 返回插件的名称，不同 ComponentPlugin 的名称不能相同
		Name() string

		// MustInitComponent 初始化 component，主要是针对一些 controller、resource 无依赖，
		// 无法在启动阶段初始化的组件，如消息、定时任务等，如果初始化失败需要 panic
		MustInitComponent()
	}
)

var (
	componentPlugins = map[string]ComponentPlugin{}
)

// RegisterComponentPlugin registers component plugin
func RegisterComponentPlugin(p ComponentPlugin) {
	if p.Name() == "" {
		panic(fmt.Errorf("%T: empty name", p))
	}

	existedPlugin, existed := componentPlugins[p.Name()]
	if existed {
		panic(fmt.Errorf("%T and %T got same name: %s", p, existedPlugin, p.Name()))
	}

	componentPlugins[p.Name()] = p
}

// MustInitComponents 初始化已注册的 Component，如果有组件初始化失败则 panic
func MustInitComponents() {
	for n, p := range componentPlugins {
		p.MustInitComponent()
		log.Infof("Init component: plugin=%s", n)
	}
}
