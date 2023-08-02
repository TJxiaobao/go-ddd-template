package entity

import "github.com/TJxiaobao/go-ddd-template/app/biz/domain/vo"

type Test struct {
	id         vo.ID
	createTime vo.Time
	context    string
	testName   vo.Username
}

func (t *Test) ID() vo.ID {
	return t.id
}

func (t *Test) CreateTime() vo.Time {
	return t.createTime
}

func (t *Test) Context() string {
	return t.context
}

func (t *Test) TestName() vo.Username {
	return t.testName
}
