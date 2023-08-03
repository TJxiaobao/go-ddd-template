package init

import (
	"github.com/TJxiaobao/go-ddd-template/app/internal/resource"
	"github.com/TJxiaobao/go-ddd-template/pkg/manager"
)

func init() {
	manager.RegisterResourcePlugin(&MySqlPlugin{})
	manager.RegisterResourcePlugin(&RedisCachePlugin{})
}

// MySqlPlugin 默认Mysql插件
type MySqlPlugin struct {
}

func (p *MySqlPlugin) Name() string {
	return "mysqlResourcePlugin"
}

func (p *MySqlPlugin) MustCreateResource() manager.Resource {
	return resource.DefaultMysqlResource()
}

// RedisCachePlugin Redis缓存插件
type RedisCachePlugin struct {
}

func (p *RedisCachePlugin) Name() string {
	return "redisCacheResourcePlugin"
}

func (p *RedisCachePlugin) MustCreateResource() manager.Resource {
	return resource.DefaultRedisCacheResource()
}
