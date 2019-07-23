package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env     string
	Port    string
	Mongodb DatabaseConfig
	Oracle  DatabaseConfig
}

type DatabaseConfig struct {
	Connection string
}

var conf = &Config{}

func Setup() *Config {
	var err error

	viper.SetDefault("ENV_STAGE", "dev")
	viper.AutomaticEnv()

	if viper.Get("env_stage") == "dev" {
		viper.AddConfigPath("./config")
		viper.SetConfigName("development")

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
