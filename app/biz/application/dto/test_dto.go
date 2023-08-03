package dto

import "github.com/TJxiaobao/go-ddd-template/app/biz/domain/entity"

type TestDto struct {
	TestId      int64  `json:"test_id"`
	TestContext string `json:"test_context"`
	TestName    string `json:"test_name"`
}

func TestToDto(test *entity.Test) *TestDto {
	return &TestDto{
		TestId:      test.ID().GetId(),
		TestContext: test.Context(),
		TestName:    test.TestName().GetUsername(),
	}
}

func TestToDtos(tests []*entity.Test) []*TestDto {
	newTests := make([]*TestDto, len(tests))
	for i := range tests {
		newTests[i] = TestToDto(tests[i])
	}
	return newTests
}
