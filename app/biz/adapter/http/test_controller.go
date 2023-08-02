package http

import (
	"context"
	"github.com/TJxiaobao/go-ddd-template/app/biz/application/app"
	"github.com/TJxiaobao/go-ddd-template/app/biz/application/cqe"
	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"github.com/TJxiaobao/go-ddd-template/pkg/manager"
	"github.com/TJxiaobao/go-ddd-template/pkg/restapi"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	testControllerOnce      sync.Once
	singletonTestController TestController
)

type TestControllerPlugin struct {
}

func (p *TestControllerPlugin) Name() string {
	return "IssueControllerPlugin"
}

func (p *TestControllerPlugin) MustCreateController() manager.Controller {
	return DefaultTestController()
}

func DefaultTestController() TestController {
	assert.NotCircular()
	testControllerOnce.Do(func() {
		singletonTestController = &testControllerImpl{
			testApp: app.DefaultTestApp(),
		}
	})
	assert.NotNil(singletonTestController)
	return singletonTestController
}

type TestController interface {
	manager.Controller
	Test(c *gin.Context)
}

type testControllerImpl struct {
	testApp app.TestApp
}

func (ctrl *testControllerImpl) RegisterOpenApi(group *gin.RouterGroup) {
	g := group.Group("/test")
	{
		g.GET("/get_list", ctrl.GetList)
		g.POST("/create", ctrl.Create)
	}
}

func (ctrl *testControllerImpl) RegisterInnerApi(group *gin.RouterGroup) {
}

func (ctrl *testControllerImpl) RegisterDebugApi(group *gin.RouterGroup) {
}

// 实现业务Handler方法

func (ctrl *testControllerImpl) GetList(c *gin.Context) {
	// todo
	query := cqe.GetTestQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		restapi.Failed(c, errno.NewSimpleError(errno.ErrParameterInvalid, err, "query"))
		return
	}

	resp, err := ctrl.testApp.GetIssues(context.Background(), &query)
	if err != nil {
		restapi.Failed(c, err)
		return
	}
	restapi.Success(c, resp)
}

func (ctrl *testControllerImpl) Create(c *gin.Context) {
	// todo
}
