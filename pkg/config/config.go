package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host   string `mapstructure:"HOST"`
	Port   int    `mapstructure:"PORT"`
	TestDB string `mapstructure:"TEST_DB"`
}

func loadConfigs(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}

func getConfigs() Config {
	configs, err := loadConfigs(".") // path: root folder
	if err != nil {
		log.Fatal("Unable to load configs: ", err.Error())
	}
	return configs
}

var DefaultConfig = getConfigs()
