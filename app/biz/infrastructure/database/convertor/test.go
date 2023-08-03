package convertor

import (
	"github.com/TJxiaobao/go-ddd-template/app/biz/domain/entity"
	"github.com/TJxiaobao/go-ddd-template/app/biz/infrastructure/database/po"
)

func ToTestEntity(testPo *po.TestModel) *entity.Test {
	return entity.NewTest(
		testPo.Id,
		testPo.CreatedAt,
		testPo.Context,
		testPo.TestName,
	)
}

func ToIssuePo(test *entity.Test) *po.TestModel {
	testPo := &po.TestModel{
		BaseModel: po.BaseModel{CreatedAt: test.CreateTime().GetTime(), Id: test.ID().GetId()},
		Context:   test.Context(),
		TestName:  test.TestName().GetUsername(),
	}
	return testPo
}
