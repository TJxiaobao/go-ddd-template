package cqe

import "github.com/TJxiaobao/go-ddd-template/pkg/errno"

type GetTestQuery struct {
	TestId    string `json:"test_id"`
	PageNum   int    `form:"page_num"`
	PageSize  int    `form:"page_size"`
	StartTime int64  `form:"start_time"`
	EndTime   int64  `form:"end_time"`
	TestName  string `json:"test_name"`
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
