package manager

import (
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	log "github.com/sirupsen/logrus"

	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
	appConfig "github.com/TJxiaobao/go-ddd-template/pkg/config"
)

type (
	DubboConsumerPlugin interface {
		// Name 返回插件的名称，不同 DubboConsumerPlugin 的名称不能相同
		Name() string

		// CreateConsumers 创建 Dubbo consumer
		CreateConsumers() []common.RPCService

		// BuildReferenceConfig 构建一个 consumer reference 配置，当需要定制协议、超时时间，泛化调用等可实现该接口。
		// 如果返回 nil，则使用默认的 tri 协议，3 秒超时，从 pb 读取接口名
		BuildReferenceConfig() *config.ReferenceConfig
	}

	DubboProviderPlugin interface {
		// Name 返回插件的名称，不同 DubboConsumerPlugin 的名称不能相同
		Name() string

		// MustCreateProvider 创建 Dubbo provider，如果创建失败需要 panic
		MustCreateProvider() common.RPCService

		// BuildServiceConfig 构建一个 provider service config，一般不需要实现，返回 nil 即可
		BuildServiceConfig() *config.ServiceConfig
	}
)

var (
	dubboConsumerPlugins   = map[string]DubboConsumerPlugin{}
	dubboProviderPlugins   = map[string]DubboProviderPlugin{}
	defaultReferenceConfig = config.NewReferenceConfigBuilder().SetProtocol("tri").Build()
	defaultServiceConfig   = config.NewServiceConfigBuilder().Build()
)

// RegisterDubboConsumerPlugin registers Dubbo consumer plugin
func RegisterDubboConsumerPlugin(p DubboConsumerPlugin) {
	if p.Name() == "" {
		panic(fmt.Errorf("%T: empty name", p))
	}

	existedPlugin, existed := dubboConsumerPlugins[p.Name()]
	if existed {
		panic(fmt.Errorf("%T and %T got same name: %s", p, existedPlugin, p.Name()))
	}

	dubboConsumerPlugins[p.Name()] = p
}

// RegisterDubboProviderPlugin registers Dubbo provider plugin
func RegisterDubboProviderPlugin(p DubboProviderPlugin) {
	if p.Name() == "" {
		panic(fmt.Errorf("%T: empty name", p))
	}

	existedPlugin, existed := dubboProviderPlugins[p.Name()]
	if existed {
		panic(fmt.Errorf("%T and %T got same name: %s", p, existedPlugin, p.Name()))
	}

	dubboProviderPlugins[p.Name()] = p
}

// MustInitDubbo 初始化已注册的 Dubbo consumer 和 provider 组件，如果初始化失败则 panic
func MustInitDubbo() {
	if len(dubboConsumerPlugins) == 0 && len(dubboProviderPlugins) == 0 {
		panic(fmt.Errorf("no Dubbo consumer and provider plugin registered"))
	}

	rootConfig := config.RootConfig{}
	rootConfigBuilder := config.NewRootConfigBuilder()
	err := appConfig.GetConfig("dubbo", &rootConfig)
	assert.Nil(err)
	if rootConfig.Application == nil || rootConfig.Protocols == nil || rootConfig.Registries == nil {
		panic(fmt.Errorf("dubbo config application, protocols or registries is not set"))
	}
	rootConfigBuilder.SetApplication(rootConfig.Application)
	rootConfigBuilder.SetProtocols(rootConfig.Protocols)
	rootConfigBuilder.SetRegistries(rootConfig.Registries)

	registerProviders(rootConfigBuilder)

	registerConsumers(rootConfigBuilder)

	err = config.Load(config.WithRootConfig(rootConfigBuilder.Build()))
	log.Infof("init dubbo: config load result=%s", err)
	assert.Nil(err)
}

func registerProviders(builder *config.RootConfigBuilder) {
	if len(dubboProviderPlugins) == 0 {
		return
	}
	providerBuilder := config.NewProviderConfigBuilder()
	for name, plugin := range dubboProviderPlugins {
		provider := plugin.MustCreateProvider()
		config.SetProviderService(provider)
		serviceConfig := plugin.BuildServiceConfig()
		if serviceConfig == nil {
			serviceConfig = defaultServiceConfig
		}
		providerBuilder.AddService(common.GetReference(provider), serviceConfig)
		log.Infof("Register Dubbo provider: plugin=%s, provider=%+v", name, provider)
	}
	builder.SetProvider(providerBuilder.Build())
}

func registerConsumers(builder *config.RootConfigBuilder) {
	if len(dubboConsumerPlugins) == 0 {
		return
	}
	consumerBuilder := config.NewConsumerConfigBuilder()
	for name, plugin := range dubboConsumerPlugins {
		consumers := plugin.CreateConsumers()
		referenceConfig := plugin.BuildReferenceConfig()
		if referenceConfig == nil {
			referenceConfig = defaultReferenceConfig
		}
		for _, consumer := range consumers {
			config.SetConsumerService(consumer)
			consumerBuilder.AddReference(common.GetReference(consumer), referenceConfig)
			log.Infof("Register Dubbo consumer: plugin=%s, consumer=%+v", name, consumer)
		}
	}
	builder.SetConsumer(consumerBuilder.Build())
}
