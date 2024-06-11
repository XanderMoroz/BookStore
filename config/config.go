package config

import (
	"log"

	// Библиотека для работы с переменными окружения
	"github.com/spf13/viper"
)

// Формируем структуру конфигурации
type Env struct {
	// ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	AppEnv                string `mapstructure:"APP_ENV"`
	AppPort               string `mapstructure:"APP_PORT"`
	ServerAddress         string `mapstructure:"SERVER_ADDRESS"`
	Dbdriver              string `mapstructure:"DB_DRIVER"`
	DBHost                string `mapstructure:"MYSQL_HOST"`
	DBPort                string `mapstructure:"MYSQL_PORT"`
	DBUser                string `mapstructure:"MYSQL_USER"`
	DBPass                string `mapstructure:"MYSQL_PASSWORD"`
	DBName                string `mapstructure:"MYSQL_DATABASE"`
	AccessTokenSecret     string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	// RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	// RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {

	env := Env{}

	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		panic(err)
	}

	if env.AppEnv == "DEV" {
		log.Println("The App is running in development env")
	}

	return &env
}
