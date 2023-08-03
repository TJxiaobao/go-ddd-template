package entity

import (
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/vo"
	"time"
)

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

func NewTest(id int64, createTime time.Time, context, testName string) *Test {
	return &Test{
		id:         vo.NewID(id),
		createTime: vo.NewTime(createTime),
		context:    context,
		testName:   vo.NewUsername(testName),
	}
}
