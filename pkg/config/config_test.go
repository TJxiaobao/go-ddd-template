package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	MustInit("testdata/test.yaml")

	// 应用基础配置
	assert.Equal(t, "demo", AppName())
	assert.Equal(t, "dev", AppEnv())
	assert.Equal(t, 8080, AppPort())

	var port int
	err := GetConfig("port", &port)
	assert.Nil(t, err)
	assert.Equal(t, 80, port)
}
