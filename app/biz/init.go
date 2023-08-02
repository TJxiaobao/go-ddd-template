package biz

import (
	"github.com/TJxiaobao/go-ddd-template/app/biz/adapter/http"
	"github.com/TJxiaobao/go-ddd-template/pkg/manager"
)

func init() {
	manager.RegisterControllerPlugin(&http.TestControllerPlugin{})
}
