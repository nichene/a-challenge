package main

import "github.com/spf13/viper"

type Config struct {
	Port string
}

func loadEnvVars() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// Defaults
	viper.SetDefault("PORT", "8080")

	return &Config{
		Port: viper.GetString("PORT"),
	}

}
