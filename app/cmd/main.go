package main

import (
	"fmt"
	"github.com/TJxiaobao/go-ddd-template/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	_ "github.com/TJxiaobao/go-ddd-template/app/internal/init"
	"github.com/TJxiaobao/go-ddd-template/pkg/config"
	"github.com/TJxiaobao/go-ddd-template/pkg/logger"
	"github.com/TJxiaobao/go-ddd-template/pkg/manager"
	"github.com/TJxiaobao/go-ddd-template/pkg/option"
)

func main() {
	opt := option.New()
	if message := opt.ShowMessage(); len(message) > 0 {
		fmt.Print(message)
		return
	}

	config.MustInit(*opt.ConfigFile)
	if *opt.ShowConfig {
		fmt.Print(config.DumpConfig())
		return
	}

	logger.MustInit()

	manager.MustInitResources()
	defer manager.CloseResources()

	manager.MustInitComponents()
	//manager.MustInitDubbo()

	route := initRoute()
	manager.MustInitControllers(
		route.Group("/api/open/xiaobao/v1"),
		route.Group("/api/inner/xiaobao/v1"),
		route.Group("/api/debug/xiaobao/v1"),
	)

	err := startRoute(route)
	if err != nil {
		log.Fatalf("start app %s failed: %v", config.AppName(), err)
	}
}

func initRoute() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.RequestId)
	route.Use(middleware.Logging())
	route.Use(middleware.Recovery())
	route.NoRoute(middleware.NotFound)
	return route
}

func startRoute(route *gin.Engine) error {
	// Debug服务
	go func() {
		debugRoute := gin.Default()
		debugRoute.GET("/metrics", gin.WrapH(promhttp.Handler()))
		_ = debugRoute.Run(fmt.Sprintf(":%d", config.DebugPort()))
	}()
	// App服务
	return route.Run(fmt.Sprintf(":%d", config.AppPort()))
}
