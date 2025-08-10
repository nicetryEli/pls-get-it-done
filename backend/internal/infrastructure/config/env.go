package config

import (
	"os"
	"sync"
)

type Environment struct {
	POSTGRES_HOST         string
	POSTGRES_PORT         string
	POSTGRES_USER         string
	POSTGRES_PASSWORD     string
	POSTGRES_TZ           string
	POSTGRES_DB           string
	POSTGRES_SSL_MODE     string
	ENVIRONMENT           string
	APP_NAME              string
	ACCESS_TOKEN_SECRET   string
	REFRESH_TOKEN_SECRET  string
	VERIFY_TOKEN_SECRET   string
	PASSWORD_TOKEN_SECRET string
	MINIO_DOMAIN          string
	MINIO_USE_SSL         bool
	MINIO_ROOT_USER       string
	MINIO_ROOT_PASSWORD   string
}

var (
	Env     *Environment
	envOnce sync.Once
)

func init() {
	envOnce.Do(func() {
		Env = &Environment{
			POSTGRES_HOST:         os.Getenv("POSTGRES_HOST"),
			POSTGRES_PORT:         os.Getenv("POSTGRES_PORT"),
			POSTGRES_USER:         os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD:     os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_TZ:           os.Getenv("POSTGRES_TZ"),
			POSTGRES_DB:           os.Getenv("POSTGRES_DB"),
			POSTGRES_SSL_MODE:     os.Getenv("POSTGRES_SSL_MODE"),
			ENVIRONMENT:           os.Getenv("ENVIRONMENT"),
			APP_NAME:              os.Getenv("APP_NAME"),
			ACCESS_TOKEN_SECRET:   os.Getenv("ACCESS_TOKEN_SECRET"),
			REFRESH_TOKEN_SECRET:  os.Getenv("REFRESH_TOKEN_SECRET"),
			VERIFY_TOKEN_SECRET:   os.Getenv("VERIFY_TOKEN_SECRET"),
			PASSWORD_TOKEN_SECRET: os.Getenv("PASSWORD_TOKEN_SECRET"),
			MINIO_DOMAIN:          os.Getenv("MINIO_DOMAIN"),
			MINIO_USE_SSL:         os.Getenv("MINIO_USE_SSL") == "true",
			MINIO_ROOT_USER:       os.Getenv("MINIO_ROOT_USER"),
			MINIO_ROOT_PASSWORD:   os.Getenv("MINIO_ROOT_PASSWORD"),
		}
	})
}
