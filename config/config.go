package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Http  HttpConfig  `mapstructure:"http"`
	Mysql MysqlConfig `mapstructure:"mysql"`
	Mongodb  MongodbConfig  `mapstructure:"mongodb"`
	AliyunOss AliyunOssConfig  `mapstructure:"aliyunOss"`
}

type HttpConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	RunModel string `mapstructure:"run_model"`
	Version  string `mapstructure:"version"`
}

type MysqlConfig struct {
	URL         string `mapstructure:"url"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxOpen     int    `mapstructure:"max_open"`
	MaxLeftTime int    `mapstructure:"max_leftTime"`
	Debug       bool   `mapstructure:"debug"`
}

type MongodbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`	//用户
	Password string `mapstructure:"password"`	//密码
	Database string `mapstructure:"database"`	//数据库名
}

type AliyunOssConfig struct {
	Endpoint string
	AccessId string
	AccessSecret string
	BucketName string
}

func NewConfig(path string) *Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Could not unmarshal config: %s", err))
	}
	return &conf
}

var WireSet = wire.NewSet(
	NewConfig,
)
