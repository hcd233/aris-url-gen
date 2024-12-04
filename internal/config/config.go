// Package config provides the configuration
package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (

	// ReadTimeout time Gin读取超时时间
	//	@update 2024-06-22 08:59:40
	ReadTimeout time.Duration

	// WriteTimeout time Gin写入超时时间
	//	@update 2024-06-22 08:59:37
	WriteTimeout time.Duration

	// Concurrency int Fiber并发数
	//	@update 2024-12-04 10:00:00
	Concurrency int

	// DomainName string 域名
	//	@update 2024-12-04 10:00:00
	DomainName string

	// MaxHeaderBytes int Gin最大头部字节数
	//	@update 2024-06-22 08:59:34
	MaxHeaderBytes int

	// LogLevel string 日志级别
	//	@update 2024-06-22 08:59:29
	LogLevel string

	// LogDirPath string 日志目录路径
	//	@update 2024-06-22 08:59:26
	LogDirPath string

	// MysqlUser string Mysql用户名
	//	@update 2024-06-22 09:00:30
	MysqlUser string

	// MysqlPassword string Mysql密码
	//	@update 2024-06-22 09:00:45
	MysqlPassword string

	// MysqlHost string Mysql主机
	//	@update 2024-06-22 09:01:02
	MysqlHost string

	// MysqlPort string Mysql端口
	//	@update 2024-06-22 09:01:18
	MysqlPort string

	// MysqlDatabase string Mysql数据库
	//	@update 2024-06-22 09:01:34
	MysqlDatabase string

	// RedisHost string Redis主机
	RedisHost string

	// RedisPort string Redis端口
	RedisPort string

	// RedisPassword string Redis密码
	RedisPassword string
)

func init() {
	initEnvironment()
}

func initEnvironment() {
	config := viper.New()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config.SetDefault("read.timeout", 10)
	config.SetDefault("write.timeout", 10)
	config.SetDefault("max.header.bytes", 1<<20)
	config.SetDefault("concurrency", 16384)
	config.SetDefault("log.level", "info")
	config.SetDefault("log.dir", "./logs")

	config.AutomaticEnv()

	ReadTimeout = time.Duration(config.GetInt("read.timeout")) * time.Second
	WriteTimeout = time.Duration(config.GetInt("write.timeout")) * time.Second
	MaxHeaderBytes = config.GetInt("max.header.bytes")
	Concurrency = config.GetInt("concurrency")
	LogLevel = config.GetString("log.level")
	LogDirPath = config.GetString("log.dir")
	DomainName = config.GetString("domain.name")
	MysqlUser = config.GetString("mysql.user")
	MysqlPassword = config.GetString("mysql.password")
	MysqlHost = config.GetString("mysql.host")
	MysqlPort = config.GetString("mysql.port")
	MysqlDatabase = config.GetString("mysql.database")

	RedisHost = config.GetString("redis.host")
	RedisPort = config.GetString("redis.port")
	RedisPassword = config.GetString("redis.password")
}
