package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresqlClient     *gorm.DB
	postgresqlClientOnce sync.Once
)

func init() {
	postgresqlClientOnce.Do(func() {
		destination := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			Env.POSTGRES_HOST,
			Env.POSTGRES_PORT,
			Env.POSTGRES_USER,
			Env.POSTGRES_PASSWORD,
			Env.POSTGRES_DB,
			Env.POSTGRES_SSL_MODE,
			Env.POSTGRES_TZ,
		)
		connection, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
			return
		}
		sqlDb, err := connection.DB()
		if err != nil {
			log.Fatalln(err)
			return
		}
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(10)
		sqlDb.SetConnMaxLifetime(30 * time.Minute)
		sqlDb.SetConnMaxIdleTime(10 * time.Minute)
		PostgresqlClient = connection
	})
}

func ClosePostgresqlClient() {
	if PostgresqlClient != nil {
		sqlDb, err := PostgresqlClient.DB()
		if err != nil {
			log.Println(err)
			return
		}
		sqlDb.Close()
	}
}
