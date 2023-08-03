package dao

import (
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"github.com/TJxiaobao/go-ddd-template/pkg/repository"
	log "github.com/sirupsen/logrus"
)

type baseDao struct {
	db     *repository.Database
	logger *log.Entry
}

func newBaseDao(repo *repository.Database) *baseDao {
	return &baseDao{
		db:     repo,
		logger: log.WithField("type", "database"),
	}
}

func (d *baseDao) Error(err error) error {
	return errno.NewError(errno.ErrDatabase, err)
}
