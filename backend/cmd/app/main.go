package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/little-tonii/gofiber-base/docs"
	"github.com/little-tonii/gofiber-base/internal/adapter/handler"
	"github.com/little-tonii/gofiber-base/internal/adapter/middleware"
	"github.com/little-tonii/gofiber-base/internal/adapter/router"
	"github.com/little-tonii/gofiber-base/internal/infrastructure/config"
	persistence_impl "github.com/little-tonii/gofiber-base/internal/infrastructure/persistence"
	provider_impl "github.com/little-tonii/gofiber-base/internal/infrastructure/provider"
	healthcheck_usecase "github.com/little-tonii/gofiber-base/internal/usecase/healthcheck"
	"github.com/segmentio/kafka-go"
)

// @title			swagger api
// @version			1.0
// @description		the api documentation for server

// @contact.nam		Tony Silvertongue
// @contact.url		https://github.com/little-tonii
// @contact.email	khuongle.workonly@gmail.com

// @host		localhost:8000
// @BasePath	/api

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					bearer token for authentication

// @securitydefinitions.oauth2.password		OAuth2Password
// @tokenUrl								http://localhost:8000/api/auth/token
// @scope.read								Grants read access
// @scope.write								Grants write access
// @scope.admin								Grants read and write access to administrative information
func main() {
	// ---------- dependency injection ----------
	log.Println("🔧 initializing configuration")
	postgresTxProvider := provider_impl.NewTransactionProviderImpl(config.PostgresqlClient)
	minioProvider := provider_impl.NewFilestoreProviderImpl(config.MinioClient)
	cacheProvider := provider_impl.NewCacheProviderImpl(config.RedisClient)
	kafkaProvider := provider_impl.NewKafkaProviderImpl(config.KafkaProducer)
	userPersis := persistence_impl.NewUserPersistenceImpl(config.PostgresqlClient)
	healthcheckUsecase := healthcheck_usecase.NewHealthcheckUsecaseImpl(
		postgresTxProvider,
		minioProvider,
		cacheProvider,
		kafkaProvider,
	)
	healthcheckHandler := handler.NewHealthcheckHandler(healthcheckUsecase)

	// ---------- fiber app setup ----------
	log.Println("🔧 setting up the server")
	app := fiber.New(fiber.Config{
		AppName:           config.Env.SERVER_NAME,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		IdleTimeout:       30 * time.Minute,
		ErrorHandler:      middleware.ErrorHandler(),
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		RequestMethods:    fiber.DefaultMethods,
	})
	app.Use(middleware.LoggingMiddleware(&middleware.LoggingConfig{Logger: config.Logger}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: false,
		AllowHeaders:     "",
	}))
	app.Use(helmet.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(middleware.Timeout(&middleware.TimeoutConfig{
		ProcessTimeout: 10 * time.Second,
		SkipPaths: []string{
			"/api/healthcheck",
		},
	}))
	app.Use(middleware.AuthGuard(&middleware.AuthGuardConfig{
		UserPersis: userPersis,
		AllowPaths: []string{
			"/api/healthcheck",
			"/api/auth/token",
			"/api/auth/register",
			"/api/auth/verify",
			"/swagger/**",
		},
		Secrets: map[string]string{
			middleware.ACCESS_TOKEN_SECRET:   config.Env.ACCESS_TOKEN_SECRET,
			middleware.VERIFY_TOKEN_SECRET:   config.Env.VERIFY_TOKEN_SECRET,
			middleware.PASSWORD_TOKEN_SECRET: config.Env.PASSWORD_TOKEN_SECRET,
			middleware.REFRESH_TOKEN_SECRET:  config.Env.REFRESH_TOKEN_SECRET,
		},
	}))
	app.Use(swagger.New(swagger.Config{
		Next:     nil,
		BasePath: "/swagger",
		FilePath: "./docs/swagger.json",
		Path:     "/api-docs",
		Title:    fmt.Sprintf("%s api documentation", config.Env.SERVER_NAME),
		CacheAge: 3600,
	}))

	// ---------- router registration ----------
	log.Println("🔧 registering routers")
	baseGroup := app.Group("/api")
	router.RegisterHealthCheckRouter(&router.HealthCheckRouterConfig{
		BaseGroup:          baseGroup,
		HealthCheckHandler: healthcheckHandler,
	})
	app.Use(middleware.NotFoundHandler())

	// ---------- start server ----------
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Printf("🚀 http server starting on port %s", config.Env.SERVER_PORT)
		if err := app.Listen(fmt.Sprintf(":%s", config.Env.SERVER_PORT)); err != nil {
			log.Fatalln(err)
		}
		log.Println("✅ http server stopped gracefully")
	}()
	go func() {
		defer wg.Done()
		healthcheckConsumer := healthcheck_usecase.NewHealthcheckConsumer(kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{
				fmt.Sprintf("%s:%s", config.Env.KAFKA_NODE_0_HOST, config.Env.KAFKA_NODE_0_PORT),
			},
			GroupID:  "healthcheck-consumer",
			Topic:    "healthcheck",
			MaxBytes: 10e6,
		}), config.Logger)
		log.Println("🚀 healthcheck consumer starting consuming")
		healthcheckConsumer.Start(ctx)
		log.Println("✅ healthcheck consumer stopped gracefully")
	}()

	// ---------- graceful shutdown ----------
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	close(shutdown)
	log.Println("⚠️ shutdown signal received, shutting down http server...")
	if err := app.ShutdownWithTimeout(time.Minute); err != nil {
		log.Println(err)
	}
	log.Println("⚠️ shutdown signal received, stopping consumers...")
	cancel()
	wg.Wait()

	// ---------- clean up resource ----------
	config.ClosePostgresqlClient()
	log.Println("✅ postgresql client connection closed")
	config.CloseRedisClient()
	log.Println("✅ redis client connection closed")
}
