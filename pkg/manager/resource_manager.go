package manager

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type (
	ResourcePlugin interface {
		// Name 返回插件的名称，不同 ResourcePlugin 的名称不能相同
		Name() string

		// MustCreateResource 创建 Resource，如果创建失败需要 panic
		MustCreateResource() Resource
	}

	Resource interface {
		// MustOpen 打开资源，如果打开失败需要 panic
		MustOpen()

		// Close 关闭资源
		Close()
	}
)

var (
	resourcePlugins = map[string]ResourcePlugin{}
	resources       []Resource
)

// RegisterResourcePlugin registers resource plugin
func RegisterResourcePlugin(p ResourcePlugin) {
	if p.Name() == "" {
		panic(fmt.Errorf("%T: empty name", p))
	}

	existedPlugin, existed := resourcePlugins[p.Name()]
	if existed {
		panic(fmt.Errorf("%T and %T got same name: %s", p, existedPlugin, p.Name()))
	}

	resourcePlugins[p.Name()] = p
}

// MustInitResources 初始化已注册的 Resource，如果失败则 panic
func MustInitResources() {
	for n, p := range resourcePlugins {
		resource := p.MustCreateResource()
		resource.MustOpen()
		resources = append(resources, resource)
		log.Infof("Init resource, plugin=%s, resource=%+v", n, resource)
	}
}

// CloseResources 关闭所有注册的资源
func CloseResources() {
	for _, resource := range resources {
		resource.Close()
		log.Infof("Close resource, resource=%+v", resource)
	}
}
