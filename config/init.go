package config

import "github.com/spf13/viper"

var AppConfig = Config{}

func Init() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(err)
	}
}
