package config

import (
	"fmt"
	"os"

	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
)

const (
	AppEnvKey = "APP_ENV"
)

const (
	EnvLocal = "local"
	EnvDebug = "debug"
	EnvPre   = "pre"
	EnvProd  = "prod"
)

var (
	configEngine Engine
	appConfig    AppConfig
)

// AppConfig 应用基础配置
type AppConfig struct {
	AppName   string `json:"name" yaml:"name"`
	AppEnv    string `json:"env" yaml:"env"`
	AppPort   int    `json:"port" yaml:"port"`
	DebugPort int    `json:"debug_port" yaml:"debug_port"`
}

func MustInit(configFile string) {
	file, err := getConfigFile(configFile)
	assert.Nil(err)

	configEngine, err = NewEngine(file)
	assert.Nil(err)

	err = GetConfig("app", &appConfig)
	if err != nil {
		fmt.Println("No config for app, so AppName(), AppEnv() and AppPort() in config is not worked")
	}
}

func getConfigFile(configFile string) (string, error) {
	if len(configFile) > 0 {
		return configFile, nil
	}
	appEnv := os.Getenv(AppEnvKey)
	if len(appEnv) == 0 {
		return "", fmt.Errorf("environment variable is not existed: %s", AppEnvKey)
	}
	return fmt.Sprintf("./conf/%s/config.yaml", appEnv), nil
}

// GetConfig 获取配置， key 为要获取的配置在配置文件中的 key，config 为获取到的配置反序列化后存储的对象
func GetConfig(key string, config interface{}) error {
	return configEngine.GetConfig(key, config)
}

// DumpConfig 将解析出的配置以 yaml 格式输出
func DumpConfig() string {
	return configEngine.DumpConfig()
}

func AppName() string {
	return appConfig.AppName
}

func AppEnv() string {
	return appConfig.AppEnv
}

func AppPort() int {
	return appConfig.AppPort
}

func DebugPort() int {
	return appConfig.DebugPort
}
