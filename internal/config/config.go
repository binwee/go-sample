package config

import "github.com/spf13/viper"

type AppConfig struct {
	Server Server
}

type Server struct {
	Port int
}

type Mysql struct {
	Username string
	Password string
	Ip       string
	Port     int
	DbName   string
	Params   string
}

func GetConfig(configFilePath string) (AppConfig, error) {
	var cfg AppConfig
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
