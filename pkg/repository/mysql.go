package repository

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database 便于多个数据源扩展
type Database struct {
	Self *gorm.DB
}

type MysqlOption struct {
	Host         string         `json:"host" yaml:"host"`
	Port         int            `json:"port" yaml:"port"`
	User         string         `json:"user" yaml:"user"`
	Password     string         `json:"password" yaml:"password"`
	DB           string         `json:"db" yaml:"db"`
	Charset      string         `json:"charset" yaml:"charset"`
	MaxOpenCount int            `json:"max_open_count" yaml:"max_open_count"`
	MaxIdleCount int            `json:"max_idle_count" yaml:"max_idle_count"`
	Log          MysqlLogConfig `json:"log" yaml:"log"`
}

type MysqlLogConfig struct {
	Path       string `json:"path" yaml:"path"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`
}

// NewDatabase 初始化数据库连接
func NewDatabase(opt *MysqlOption) (*Database, error) {
	conn, err := newMysqlConn(opt)
	if err != nil {
		return nil, err
	}
	return &Database{Self: conn}, nil
}

// NewDatabaseTx 用给定的tx创建Database对象
func NewDatabaseTx(tx *gorm.DB) *Database {
	return &Database{Self: tx}
}

// Close 关闭数据库连接
func (db *Database) Close() {
	if sqlDB, err := db.Self.DB(); err == nil {
		_ = sqlDB.Close()
	}
}

func newMysqlConn(opt *MysqlOption) (*gorm.DB, error) {
	if opt.Port == 0 {
		opt.Port = 3306
	}
	if opt.Charset == "" {
		opt.Charset = "utf8mb4"
	}
	if opt.Log.MaxSize == 0 {
		opt.Log.MaxSize = 100
	}
	if opt.Log.MaxBackups == 0 {
		opt.Log.MaxBackups = 3
	}

	loggerWriter := log.New()
	if opt.Log.Path != "" {
		loggerWriter.SetOutput(&lumberjack.Logger{
			Filename:   opt.Log.Path,
			MaxSize:    opt.Log.MaxSize,
			MaxBackups: opt.Log.MaxBackups,
		})
	}
	gormLogger := logger.New(loggerWriter, logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,       // Disable color
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=5s&parseTime=%t&loc=%s",
		opt.User, opt.Password, opt.Host, opt.Port, opt.DB, opt.Charset, true, "Local")

	gormDB, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		CreateBatchSize:        100,
		SkipDefaultTransaction: false,
		Logger:                 gormLogger,
	})
	if err != nil {
		return nil, err
	}

	// set for db connection
	if sqlDB, err := gormDB.DB(); err == nil {
		// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
		sqlDB.SetMaxOpenConns(opt.MaxOpenCount)
		// 用于设置闲置的连接数，设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
		sqlDB.SetMaxIdleConns(opt.MaxIdleCount)
		// 用于设置一个连接可被复用的最大时长。
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return gormDB, nil
}
