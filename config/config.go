package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
) // Библиотека для работы с переменными окружения

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
	// viper.AddConfigPath("./envs")
	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}

// Формируем структуру конфигурации для подключения к базе данных
type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

type AppEnvConfig struct {
	AppSecret  any
	Dbdriver   any
	DbUser     any
	DbPassword any
	DbHost     any
	DbPort     any
	DbName     any
}

func LoadConfig() *AppEnvConfig {
	viper.AddConfigPath("./envs")
	viper.SetConfigName("dev")
	// viper.SetConfigType("env")
	viper.SetConfigFile("./envs/.env")

	// viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println("ошибка извлечения переменных окружения:", err)
		panic(err)
	}
	fmt.Println(viper.Get("PORT"))

	config := &AppEnvConfig{
		AppSecret:  viper.Get("APP_SECRET"),
		Dbdriver:   viper.Get("DB_DRIVER"),
		DbUser:     viper.Get("DB_USER"),
		DbPassword: viper.Get("DB_PASSWORD"),
		DbHost:     viper.Get("DB_HOST"),
		DbPort:     viper.Get("PORT"),
		DbName:     viper.Get("DB_NAME"),
	}

	return config
}
