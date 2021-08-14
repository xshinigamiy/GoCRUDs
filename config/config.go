package config

import (
	"GoPractice/GoCRUDs/pkg/utils"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	ServiceConfig *ServiceConfigs
	DBConfigs     *DBConfigs
}

type DBConfigs struct {
	Url          string
	Port         string
	DatabaseName string
	DBDialect    string
}

type ServiceConfigs struct {
	AppName string
	AppPort string
}

func GetConfigs() *Config {
	once.Do(InitConfigs)
	return config
}

func InitConfigs() {
	config = &Config{
		ServiceConfig: &ServiceConfigs{
			AppName: utils.GetEnv("APP_NAME"),
			AppPort: utils.GetEnv("APP_PORT"),
		},
		DBConfigs: &DBConfigs{
			Url:          utils.GetEnv("DB_URL"),
			Port:         utils.GetEnv("DB_PORT"),
			DBDialect:    utils.GetEnv("DB_DIALECT"),
			DatabaseName: utils.GetEnv("DB_DATABASE_NAME"),
		},
	}
}
