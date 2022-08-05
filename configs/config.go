package configs

import (
	"fmt"
	"os"
	"sync"
)

type ConfigSet struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

var lock = &sync.Mutex{}
var appConfig *ConfigSet

func InitConfig() *ConfigSet {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func Get() *ConfigSet {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *ConfigSet {
	var config ConfigSet
	config.Host = getEnv("DB_HOST", "")
	config.Port = getEnv("DB_PORT", "")
	config.Username = getEnv("DB_USERNAME", "")
	config.Password = getEnv("DB_PASSWORD", "")
	config.Name = getEnv("DB_NAME", "")

	// Info
	fmt.Println(config)

	return &config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
