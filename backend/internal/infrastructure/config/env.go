package config

import (
	"os"
	"sync"
)

type Environment struct {
	POSTGRES_HOST             string
	POSTGRES_PORT             string
	POSTGRES_USER             string
	POSTGRES_PASSWORD         string
	POSTGRES_SESSION_TZ       string
	POSTGRES_DB               string
	POSTGRES_SESSION_SSL_MODE string

	ENVIRONMENT           string
	SERVER_NAME           string
	SERVER_PORT           string
	ACCESS_TOKEN_SECRET   string
	REFRESH_TOKEN_SECRET  string
	VERIFY_TOKEN_SECRET   string
	PASSWORD_TOKEN_SECRET string

	MINIO_HOST            string
	MINIO_API_PORT_NUMBER string
	MINIO_USE_SSL         bool
	MINIO_ROOT_USER       string
	MINIO_ROOT_PASSWORD   string

	REDIS_HOST        string
	REDIS_PORT_NUMBER string
	REDIS_PASSWORD    string

	RABBITMQ_USERNAME         string
	RABBITMQ_PASSWORD         string
	RABBITMQ_VHOST            string
	RABBITMQ_HOST             string
	RABBITMQ_NODE_PORT_NUMBER string

	KAFKA_NODE_0_HOST string
	KAFKA_NODE_0_PORT string
}

var (
	Env     *Environment
	envOnce sync.Once
)

func init() {
	envOnce.Do(func() {
		Env = &Environment{
			POSTGRES_HOST:             os.Getenv("POSTGRES_HOST"),
			POSTGRES_PORT:             os.Getenv("POSTGRES_PORT"),
			POSTGRES_USER:             os.Getenv("POSTGRES_USER"),
			POSTGRES_PASSWORD:         os.Getenv("POSTGRES_PASSWORD"),
			POSTGRES_DB:               os.Getenv("POSTGRES_DB"),
			POSTGRES_SESSION_TZ:       os.Getenv("POSTGRES_SESSION_TZ"),
			POSTGRES_SESSION_SSL_MODE: os.Getenv("POSTGRES_SESSION_SSL_MODE"),
			ENVIRONMENT:               os.Getenv("ENVIRONMENT"),
			SERVER_NAME:               os.Getenv("SERVER_NAME"),
			ACCESS_TOKEN_SECRET:       os.Getenv("ACCESS_TOKEN_SECRET"),
			REFRESH_TOKEN_SECRET:      os.Getenv("REFRESH_TOKEN_SECRET"),
			VERIFY_TOKEN_SECRET:       os.Getenv("VERIFY_TOKEN_SECRET"),
			PASSWORD_TOKEN_SECRET:     os.Getenv("PASSWORD_TOKEN_SECRET"),
			MINIO_HOST:                os.Getenv("MINIO_HOST"),
			MINIO_API_PORT_NUMBER:     os.Getenv("MINIO_API_PORT_NUMBER"),
			MINIO_USE_SSL:             os.Getenv("MINIO_USE_SSL") == "true",
			MINIO_ROOT_USER:           os.Getenv("MINIO_ROOT_USER"),
			MINIO_ROOT_PASSWORD:       os.Getenv("MINIO_ROOT_PASSWORD"),
			REDIS_HOST:                os.Getenv("REDIS_HOST"),
			REDIS_PORT_NUMBER:         os.Getenv("REDIS_PORT_NUMBER"),
			REDIS_PASSWORD:            os.Getenv("REDIS_PASSWORD"),
			RABBITMQ_USERNAME:         os.Getenv("RABBITMQ_USERNAME"),
			RABBITMQ_PASSWORD:         os.Getenv("RABBITMQ_PASSWORD"),
			RABBITMQ_HOST:             os.Getenv("RABBITMQ_HOST"),
			RABBITMQ_NODE_PORT_NUMBER: os.Getenv("RABBITMQ_NODE_PORT_NUMBER"),
			SERVER_PORT:               os.Getenv("SERVER_PORT"),
			RABBITMQ_VHOST:            os.Getenv("RABBITMQ_VHOST"),
			KAFKA_NODE_0_HOST:         os.Getenv("KAFKA_NODE_0_HOST"),
			KAFKA_NODE_0_PORT:         os.Getenv("KAFKA_NODE_0_PORT"),
		}
	})
}
