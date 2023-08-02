package config

import (
	"errors"
	"io/fs"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEngineNotExistFile(t *testing.T) {
	engine, err := NewEngine("testdata/not_exist")
	assert.Nil(t, engine)
	assert.True(t, errors.Is(err, syscall.ENOENT))

	var pathErr *fs.PathError
	assert.True(t, errors.As(err, &pathErr))
	assert.Equal(t, syscall.ENOENT, pathErr.Err)
}

func TestNewEngineJsonFile(t *testing.T) {
	engine, err := NewEngine("testdata/test.json")
	assert.Nil(t, err)

	resultAssert(t, engine)
}

func TestNewEngineYamlFile(t *testing.T) {
	engine, err := NewEngine("testdata/test.yaml")
	assert.Nil(t, err)

	resultAssert(t, engine)
}

func TestLocalCenterJsonFile(t *testing.T) {
	engine, err := NewEngine("testdata/test_local.json")
	assert.Nil(t, err)
	var listenPort int
	engine.GetConfig("listen_port", &listenPort)
	assert.Equal(t, 8081, listenPort)
}

func TestNacosCenterJsonFile(t *testing.T) {
	engine, err := NewEngine("testdata/test_nacos.json")
	assert.Nil(t, err)
	var listenPort int
	engine.GetConfig("listen_port", &listenPort)
	assert.Equal(t, 8081, listenPort)
}

func TestNacosCenterNoGroup(t *testing.T) {
	_, err := NewEngine("testdata/test_nacos_no_group.json")
	assert.Error(t, err)
}

func resultAssert(t *testing.T, engine Engine) {
	// 单个 int 配置
	var port int
	engine.GetConfig("port", &port)
	assert.Equal(t, 80, port)

	// dict 配置
	var dict map[string]string
	engine.GetConfig("dict", &dict)
	assert.Equal(t, "def", dict["abc"])
	assert.Equal(t, "456", dict["123"])
}
