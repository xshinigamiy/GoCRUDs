package config

import (
	"GoCRUDs/pkg/utils"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	ServiceConfig *ServiceConfigs
	DBConfigs     *DBConfigs
	Env           string
}

type DBConfigs struct {
	Url          string
	Port         string
	DatabaseName string
	DBDialect    string
	DBUser       string
	DBPassword   string
}

type ServiceConfigs struct {
	AppName string
	AppPort string
}

func init() {
	InitConfigs()
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
			DBPassword:   utils.GetEnv("DB_USER"),
			DBUser:       utils.GetEnv("DB_PASSWORD"),
		},
		Env: utils.GetEnv("ENV_NAME"),
	}

	log.Info(config.DBConfigs.Url)
}
