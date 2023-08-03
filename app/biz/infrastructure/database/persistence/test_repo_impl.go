package persistence

import (
	"context"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/entity"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/repo"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/convertor"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/dao"
)

type testRepoImpl struct {
	testDao *dao.TestDao
}

func NewTestRepo(testDao *dao.TestDao) *testRepoImpl {
	return &testRepoImpl{
		testDao: testDao,
	}
}

func (r *testRepoImpl) FindByTestQuery(ctx context.Context, testQuery repo.TestQuery) ([]*entity.Test, error) {
	tests, err := r.testDao.GetByTestQuery(ctx, testQuery)
	if err != nil {
		return nil, err
	}

	res := make([]*entity.Test, len(tests))
	for i := range tests {
		res[i] = convertor.ToTestEntity(tests[i])
	}
	return res, nil
}

func (r *testRepoImpl) FindByTestId(ctx context.Context, testId string) (*entity.Test, error) {
	testModel, err := r.testDao.SelectByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}
	if testModel == nil {
		return nil, nil
	}
	return convertor.ToTestEntity(testModel), nil
}

func (r *testRepoImpl) Save(ctx context.Context, test *entity.Test) error {
	testModel := convertor.ToIssuePo(test)
	var err error
	if testModel.Id > 0 {
		err = r.testDao.Update(ctx, testModel)
	} else {
		err = r.testDao.Insert(ctx, testModel)
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *testRepoImpl) Remove(ctx context.Context, testId string) error {
	err := r.testDao.Delete(ctx, testId)
	if err != nil {
		return err
	}
	return nil
}
