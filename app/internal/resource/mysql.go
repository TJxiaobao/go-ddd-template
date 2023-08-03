package resource

import (
	"fmt"
	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
	"github.com/TJxiaobao/go-ddd-template/pkg/config"
	"github.com/TJxiaobao/go-ddd-template/pkg/repository"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	mysqlOnce              sync.Once
	singletonMysqlResource *MySqlResource
)

// MySqlResource 多数据源都注入到MySqlResource内
type MySqlResource struct {
	rwRepo *repository.Database
	roRepo *repository.Database
}

func DefaultMysqlResource() *MySqlResource {
	mysqlOnce.Do(func() {
		singletonMysqlResource = &MySqlResource{}
	})
	assert.NotNil(singletonMysqlResource)
	return singletonMysqlResource
}

func (r *MySqlResource) MustOpen() {
	if r.rwRepo == nil {
		r.rwRepo = newMySqlRepo("mysql_rw")
	}
	assert.NotNil(r.rwRepo)
}

func (r *MySqlResource) Close() {
	if r.rwRepo != nil {
		r.rwRepo.Close()
	}
	if r.roRepo != nil {
		r.roRepo.Close()
	}
}

func (r *MySqlResource) RwRepo() *repository.Database {
	return r.rwRepo
}

func (r *MySqlResource) RoRepo() *repository.Database {
	return r.roRepo
}

func newMySqlRepo(dbCfgKey string) *repository.Database {
	var (
		opt  repository.MysqlOption
		repo *repository.Database
	)

	err := config.GetConfig(dbCfgKey, &opt)
	if err != nil {
		log.Errorf("init mysql failed: %v", err)
		return nil
	}

	if opt.Log.Path == "" && config.AppEnv() != config.EnvLocal {
		opt.Log.Path = fmt.Sprintf("./home/log/%s/mysql.log", config.AppName())
	}

	repo, err = repository.NewDatabase(&opt)
	if err != nil {
		log.Errorf("init mysql failed: %v", err)
		return nil
	}

	return repo
}
