package config

import (
	"fmt"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// struct config like in config json
type AppConfig struct {
	BackendPort int `json:"backend_port"`
	Mysql       struct {
		Host              string `json:"host"`
		Port              string `json:"port"`
		User              string `json:"user"`
		Password          string `json:"password"`
		Name              string `json:"name"`
		MaxIdleConnection int    `json:"max_idle_con"`
		MaxOpenConnection int    `json:"max_open_con"`
	} `json:"mysql"`
	JWT struct {
		Secret    string `json:"secret"`
		ExpiredAt int    `json:"expired_at"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

// GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	// re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.BackendPort = 7070

	defaultConfig.Mysql.Host = "127.0.0.1"
	defaultConfig.Mysql.Port = "3306"
	defaultConfig.Mysql.User = "root"
	defaultConfig.Mysql.Password = ""
	defaultConfig.Mysql.Name = "boilerplate"
	defaultConfig.Mysql.MaxIdleConnection = 5
	defaultConfig.Mysql.MaxOpenConnection = 10

	defaultConfig.JWT.Secret = ""
	defaultConfig.JWT.ExpiredAt = 1440

	// Set the file name of the configuration file
	viper.SetConfigName("config")

	// Set the configuration file type
	viper.SetConfigType("json")

	// Set the path to look for the configuration file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	var finalConfig AppConfig

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config gile, %s", err)
	}

	err = viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Error("failed to extract config, will use default value", err)
		return &defaultConfig
	}

	return &finalConfig

}
