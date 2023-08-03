package dao

import (
	"context"
	"errors"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/repo"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/po"
	"github.com/TJxiaobao/go-ddd-template/pkg/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TestDao struct {
	*baseDao
}

func NewTestDao(repo *repository.Database) *TestDao {
	return &TestDao{newBaseDao(repo)}
}

func (d *TestDao) buildTestQueryCondition(query repo.TestQuery) *gorm.DB {
	tx := d.db.Self
	if query.StartTime > 0 {
		tx = tx.Where("gmt_create >= ?", time.Unix(query.StartTime, 0))
	}
	if query.EndTime > 0 {
		tx = tx.Where("gmt_create <= ?", time.Unix(query.EndTime, 0))
	}
	if query.Username != "" {
		tx = tx.Where("test_name = ?", query.Username)
	}
	return tx.Where("is_deleted = 0")
}

func (d *TestDao) GetByTestQuery(ctx context.Context, query repo.TestQuery) ([]*po.TestModel, error) {
	var (
		res    []*po.TestModel
		offset = (query.PageNum - 1) * query.PageSize
	)
	condition := d.buildTestQueryCondition(query)
	err := d.db.Self.Model(&po.TestModel{}).Where(condition).Offset(offset).Order("id DESC").Limit(query.PageSize).Find(&res).Offset(-1).Limit(-1).Error
	if err != nil {
		d.logger.Errorf("GetByIssueQuery error: %v", err)
		return nil, d.Error(err)
	}
	return res, nil
}

func (d *TestDao) SelectByTestId(ctx context.Context, testId string) (*po.TestModel, error) {
	var m po.TestModel
	err := d.db.Self.Where("test_id = ? AND is_deleted = 0", testId).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByTestId error: %v", err)
		return nil, d.Error(err)
	}
	return &m, nil
}

func (d *TestDao) Upsert(ctx context.Context, test *po.TestModel) error {
	err := d.db.Self.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "test_id"}},
		UpdateAll: true,
	}).Create(test).Error
	if err != nil {
		d.logger.Errorf("Upsert error: %v", err)
		return d.Error(err)
	}
	return nil
}

func (d *TestDao) Insert(ctx context.Context, test *po.TestModel) error {
	err := d.db.Self.Create(test).Error
	if err != nil {
		d.logger.Errorf("Insert error: %v", err)
		return d.Error(err)
	}
	return nil
}

func (d *TestDao) Update(ctx context.Context, test *po.TestModel) error {
	err := d.db.Self.Select("*").Omit("id, gmt_create, test_id").Where("id = ?", test.Id).Updates(test).Error
	if err != nil {
		d.logger.Errorf("Update error: %v", err)
		return d.Error(err)
	}
	return nil
}

func (d *TestDao) Delete(ctx context.Context, testId string) error {
	err := d.db.Self.Model(&po.TestModel{}).Where("id = ?", testId).Update("is_deleted", time.Now().Unix()).Error
	if err != nil {
		d.logger.Errorf("Delete error: %v", err)
		return d.Error(err)
	}
	return nil
}
