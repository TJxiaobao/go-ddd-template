package manager

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type (
	ControllerPlugin interface {
		// Name 返回插件的名称，不同 ControllerPlugin 的名称不能相同
		Name() string

		// MustCreateController 创建 Controller，如果创建失败需要 panic
		MustCreateController() Controller
	}

	Controller interface {
		// RegisterOpenApi 在给定的 group 下注册相应的 open api group 及相应的路由处理方法
		RegisterOpenApi(group *gin.RouterGroup)

		// RegisterInnerApi 在给定的 group 下注册相应的 inner api group 及相应的路由处理方法
		RegisterInnerApi(group *gin.RouterGroup)

		// RegisterDebugApi 在给定的 group 下注册相应的 debug api group 及相应的路由处理方法
		RegisterDebugApi(group *gin.RouterGroup)

		// RegisterOpsApi 在给定的 group 下注册相应的 ops api group 及相应的路由处理方法
		RegisterOpsApi(group *gin.RouterGroup)
	}
)

var (
	controllerPlugins = map[string]ControllerPlugin{}
)

// RegisterControllerPlugin registers controller plugin
func RegisterControllerPlugin(p ControllerPlugin) {
	if p.Name() == "" {
		panic(fmt.Errorf("%T: empty name", p))
	}

	existedPlugin, existed := controllerPlugins[p.Name()]
	if existed {
		panic(fmt.Errorf("%T and %T got same name: %s", p, existedPlugin, p.Name()))
	}

	controllerPlugins[p.Name()] = p
}

// MustInitControllers 初始化已注册的 Controller，包括相关 Controller 的创建及 api group 的注册，如果失败则 panic
func MustInitControllers(openApiGroup, innerApiGroup, debugApiGroup, opsApiGroup *gin.RouterGroup) {
	for n, p := range controllerPlugins {
		controller := p.MustCreateController()
		if openApiGroup != nil {
			controller.RegisterOpenApi(openApiGroup)
		}
		if innerApiGroup != nil {
			controller.RegisterInnerApi(innerApiGroup)
		}
		if debugApiGroup != nil {
			controller.RegisterDebugApi(debugApiGroup)
		}
		if opsApiGroup != nil {
			controller.RegisterOpsApi(opsApiGroup)
		}
		log.Infof("Register controller: plugin=%s, controller=%+v", n, controller)
	}
}
