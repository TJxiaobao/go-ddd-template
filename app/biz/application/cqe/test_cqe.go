package cqe

import "github.com/TJxiaobao/go-ddd-template/pkg/errno"

type GetTestQuery struct {
	TestId string `json:"test_id"`
}

func (q *GetTestQuery) Validate() error {
	if q.TestId == "" {
		return errno.NewSimpleError(errno.ErrMissingParameter, nil, "test_id")
	}
	return nil
}

type CreateTestCmd struct {
	TestName string `json:"test_name"`
	Context  string `json:"context"`
}

func (q *CreateTestCmd) Validate() error {
	if q.TestName == "" {
		return errno.NewSimpleError(errno.ErrMissingParameter, nil, "test_name")
	}
	return nil
}
