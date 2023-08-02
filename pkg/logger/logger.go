package logger

import (
	"fmt"
	"github.com/TJxiaobao/go-ddd-template/pkg/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option struct {
	Path       string `json:"path" yaml:"path"`
	LogLevel   string `json:"log_level" yaml:"log_level"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`
	OmitCaller bool   `json:"omit_caller" yaml:"omit_caller"`
}

func MustInit() {
	var opt Option
	err := config.GetConfig("logger", &opt)
	if err != nil {
		log.Warnf("logger config not found, use default")
	}

	// 本地调试环境日志不输出到文件中
	if config.AppEnv() != config.EnvLocal {
		if opt.Path == "" {
			// 当日志路径没有配置时将日志写到统一路径
			opt.Path = fmt.Sprintf("/home/log/%s/run.log", config.AppName())
		}
		if opt.MaxBackups == 0 {
			opt.MaxBackups = 3
		}
		if opt.MaxSize == 0 {
			opt.MaxSize = 100
		}
		log.SetOutput(&lumberjack.Logger{
			Filename:   opt.Path,
			MaxBackups: opt.MaxBackups,
			MaxSize:    opt.MaxSize,
		})
	}

	switch opt.LogLevel {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(!opt.OmitCaller)
	log.AddHook(&DefaultFieldHook{appName: config.AppName()})
}

type DefaultFieldHook struct {
	appName string
}

func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	entry.Data["app_name"] = hook.appName
	return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}
