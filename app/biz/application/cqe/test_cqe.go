package cqe

import "github.com/TJxiaobao/go-ddd-template/pkg/errno"

type GetTestQuery struct {
	TestId   string `json:"test_id"`
	TestName string `json:"test_name"`
}

func (q *GetTestQuery) Validate() error {
	if q.TestId == "" {
		return errno.NewSimpleError(errno.ErrMissingParameter, nil, "test_id")
	}
	return nil
}
