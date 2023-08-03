package app

import (
	"context"
	"github.com/TJxiaobao/go-ddd-template/app/biz/application/cqe"
	"github.com/TJxiaobao/go-ddd-template/app/biz/application/dto"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/repo"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/service"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/dao"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/persistence"
	"github.com/TJxiaobao/go-ddd-template/app/internal/resource"
	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
	"sync"
)

var (
	testAppOnce      sync.Once
	singletonTestApp TestApp
)

type TestApp interface {
	Create(background context.Context, c *cqe.CreateTestCmd) error
	GetList(background context.Context, query *cqe.GetTestQuery) (*dto.PageResult, error)
}

type testApp struct {
	testRepo repo.TestRepo
	testSrv  service.TestService
}

func DefaultTestApp() TestApp {
	assert.NotCircular()
	testAppOnce.Do(func() {
		var (
			db       = resource.DefaultMysqlResource().RwRepo()
			testDao  = dao.NewTestDao(db)
			testRepo = persistence.NewTestRepo(testDao)
		)
		singletonTestApp = &testApp{
			testRepo: testRepo,
			testSrv:  service.NewTestService(testRepo),
		}
	})
	assert.NotNil(singletonTestApp)
	return singletonTestApp
}

func (t *testApp) GetList(ctx context.Context, query *cqe.GetTestQuery) (*dto.PageResult, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	testQuery := repo.TestQuery{
		PageNum:   query.PageNum,
		PageSize:  query.PageSize,
		StartTime: query.StartTime,
		EndTime:   query.EndTime,
		Username:  query.TestName,
		TestId:    query.TestId,
	}
	tests, err := t.testRepo.FindByTestQuery(ctx, testQuery)
	if err != nil {
		return nil, err
	}

	testDtos := dto.TestToDtos(tests)
	return dto.NewPageResult(testDtos, query.PageNum, query.PageSize), nil
}

func (t *testApp) Create(ctx context.Context, cmd *cqe.CreateTestCmd) error {
	if err := cmd.Validate(); err != nil {
		return err
	}

	return nil
}
