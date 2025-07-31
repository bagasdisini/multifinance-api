package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	AppHost          string
	AppPort          string
	CORSAllowOrigins string

	SwaggerHost string

	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
)

func init() {
	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("CORS_ALLOW_ORIGINS", "*")

	viper.SetDefault("SWAGGER_HOST", "0.0.0.0:8080")

	viper.SetDefault("DATABASE_HOST", "localhost")
	viper.SetDefault("DATABASE_PORT", "5432")
	viper.SetDefault("DATABASE_USERNAME", "postgres")
	viper.SetDefault("DATABASE_PASSWORD", "postgres")
	viper.SetDefault("DATABASE_NAME", "multifinance")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found, using environment variables")
	}

	AppHost = viper.GetString("APP_HOST")
	AppPort = viper.GetString("APP_PORT")
	CORSAllowOrigins = viper.GetString("CORS_ALLOW_ORIGINS")

	SwaggerHost = viper.GetString("SWAGGER_HOST")

	DatabaseHost = viper.GetString("DATABASE_HOST")
	DatabasePort = viper.GetString("DATABASE_PORT")
	DatabaseUsername = viper.GetString("DATABASE_USERNAME")
	DatabasePassword = viper.GetString("DATABASE_PASSWORD")
	DatabaseName = viper.GetString("DATABASE_NAME")
}
