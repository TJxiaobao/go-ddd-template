package repo

import (
	"context"
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/entity"
)

type TestRepo interface {
	FindByTestQuery(ctx context.Context, testQuery TestQuery) ([]*entity.Test, error)
	FindByTestId(ctx context.Context, issueId string) (*entity.Test, error)
	Save(ctx context.Context, test *entity.Test) error
	Remove(ctx context.Context, testId string) error
}

type TestQuery struct {
	PageNum   int
	PageSize  int
	StartTime int64
	EndTime   int64
	Username  string
	TestId    string
}
