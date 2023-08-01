package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Engine interface {
	GetConfig(key string, config interface{}) error
	DumpConfig() string
}

type engineImpl struct {
	config map[string]interface{}
}

type configCenter struct {
	Type    string `yaml:"type"`
	Proto   string `yaml:"proto"`
	Address string `yaml:"address"`
	Config  struct {
		Port   string `yaml:"port"`
		DataId string `yaml:"data_id"`
		Group  string `yaml:"group"`
	} `yaml:"config"`
}

var (
	supportConfigTypeSuffix = map[string]struct{}{".json": {}, ".yaml": {}, ".yml": {}}
	configCenterKey         = "config_center"
)

func NewEngine(configFile string) (Engine, error) {
	data, err := readFileConfig(configFile)
	if err != nil {
		return nil, err
	}

	dataCenter, err := readCenterConfig(data)
	if err != nil {
		return nil, err
	}

	err = mergeConfig(data, dataCenter)
	if err != nil {
		return nil, err
	}

	return &engineImpl{config: data}, nil
}

func (e *engineImpl) GetConfig(key string, config interface{}) error {
	return getConfig(key, e.config, config)
}

func (e *engineImpl) DumpConfig() string {
	bytes, err := yaml.Marshal(e.config)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func readFileConfig(file string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	ext := path.Ext(file)
	if _, ok := supportConfigTypeSuffix[ext]; !ok {
		return nil, fmt.Errorf("not support config type: %s", file)
	}

	var data = map[string]interface{}{}
	err = unmarshal(ext == ".json", bytes, &data)
	return data, err
}

func readCenterConfig(data map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := data[configCenterKey]; !ok {
		return nil, nil
	}

	var centerConfig configCenter
	err := getConfig(configCenterKey, data, &centerConfig)
	if err != nil {
		return nil, err
	}

	var dataCenter map[string]interface{}
	if centerConfig.Type == EnvLocal {
		dataCenter, err = readFileConfig(centerConfig.Address)
	} else if centerConfig.Type == "nacos" {
		dataCenter, err = readNacosConfig(centerConfig)
	} else {
		err = fmt.Errorf("not support config center type: %s", centerConfig.Type)
	}
	return dataCenter, err
}

func readNacosConfig(centerConfig configCenter) (map[string]interface{}, error) {
	// 不需要 namespace 参数，测试发现随便传什么  namespace 值不影响
	url := fmt.Sprintf(
		"http://%s:%s/nacos/v1/cs/configs?dataId=%s&group=%s",
		centerConfig.Address, centerConfig.Config.Port, centerConfig.Config.DataId, centerConfig.Config.Group)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read nacos error: status=%d, body=%s", response.StatusCode, string(bytes))
	}

	var data = map[string]interface{}{}
	err = unmarshal(response.Header.Get("Config-Type") == "json", bytes, &data)
	return data, err
}

func unmarshal(isJson bool, in []byte, out interface{}) error {
	if isJson {
		return json.Unmarshal(in, out)
	}
	return yaml.Unmarshal(in, out)
}

func getConfig(key string, data map[string]interface{}, out interface{}) error {
	if len(key) == 0 {
		return fmt.Errorf("key is blank")
	}
	value, ok := data[key]
	if !ok {
		return fmt.Errorf("config for key is not existed: %s", key)
	}
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, out)
}

func mergeConfig(dist, source map[string]interface{}) error {
	for k, v := range source {
		if _, ok := dist[k]; ok {
			return fmt.Errorf("config key is existed: %s", k)
		}
		dist[k] = v
	}
	return nil
}
