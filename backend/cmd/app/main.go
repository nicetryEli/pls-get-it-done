package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
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
	healthcheck_usecase "github.com/little-tonii/gofiber-base/internal/usecase/healthcheck"
)

// @title			swagger api
// @version			1.0
// @description		the api documentation for server

// @contact.nam		Anna Lilly
// @contact.url		https://github.com/nicetryEli
// @contact.email	annalilly131205@gmail.com

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
	postgresTxProvider := persistence_impl.NewTransactionProviderImpl(config.PostgresqlClient)
	minioProvider := persistence_impl.NewFilestoreProviderImpl(config.MinioClient)
	userPersis := persistence_impl.NewUserPersistenceImpl(config.PostgresqlClient)
	healthcheckUsecase := healthcheck_usecase.NewHealthcheckUsecaseImpl(postgresTxProvider, minioProvider)
	healthcheckHandler := handler.NewHealthcheckHandler(healthcheckUsecase)

	// ---------- fiber app setup ----------
	log.Println("🔧 setting up the server")
	app := fiber.New(fiber.Config{
		AppName:           config.Env.APP_NAME,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		IdleTimeout:       30 * time.Minute,
		ErrorHandler:      middleware.ErrorHandler(),
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
		ReadTimeout:       10 * time.Second,
		RequestMethods:    fiber.DefaultMethods,
	})
	app.Use(middleware.LoggingMiddleware(config.Logger))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: false,
		AllowHeaders:     "",
	}))
	app.Use(helmet.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(middleware.AuthGuard(
		userPersis,
		[]string{
			"/api/healthcheck",
			"/api/auth/token",
			"/api/auth/register",
			"/api/auth/verify",
			"/swagger/**",
		},
		map[string]string{
			middleware.ACCESS_TOKEN_SECRET:   config.Env.ACCESS_TOKEN_SECRET,
			middleware.VERIFY_TOKEN_SECRET:   config.Env.VERIFY_TOKEN_SECRET,
			middleware.PASSWORD_TOKEN_SECRET: config.Env.PASSWORD_TOKEN_SECRET,
			middleware.REFRESH_TOKEN_SECRET:  config.Env.REFRESH_TOKEN_SECRET,
		},
	))
	app.Use(swagger.New(swagger.Config{
		Next:     nil,
		BasePath: "/swagger",
		FilePath: "./docs/swagger.json",
		Path:     "/api-docs",
		Title:    fmt.Sprintf("%s api documentation", config.Env.APP_NAME),
		CacheAge: 3600,
	}))

	// ---------- router registration ----------
	log.Println("🔧 registering routers")
	baseGroup := app.Group("/api")
	router.RegisterHealthCheckRouter(baseGroup, healthcheckHandler)
	app.Use(middleware.NotFoundHandler())

	// ---------- start server ----------
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("🚀 http server started on port 8000")

	// ---------- graceful shutdown ----------
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	close(shutdown)
	log.Println("⚠️ shutdown signal received, shutting down http server...")
	if err := app.ShutdownWithTimeout(time.Minute); err != nil {
		log.Println(err)
	}
	log.Println("✅ http server stopped gracefully.")

	// ---------- clean up resource ----------
	config.ClosePostgresqlClient()
	log.Println("✅ PostgreSQL client closed.")
}
