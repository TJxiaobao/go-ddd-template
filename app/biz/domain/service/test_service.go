package service

import (
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/repo"
)

type TestService interface {
	// 本层svr主要是用来实现两个或以上领域之间关联的业务，能不使用就不使用！
}

type testServiceImp struct {
	testRepo repo.TestRepo
}

func NewTestService(testRepo repo.TestRepo) *testServiceImp {
	return &testServiceImp{
		testRepo: testRepo,
	}
}
