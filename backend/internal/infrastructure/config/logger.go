package config

import (
	"log"
	"strings"
	"sync"

	"go.uber.org/zap"
)

var (
	Logger     *zap.Logger
	loggerOnce sync.Once
)

func init() {
	loggerOnce.Do(func() {
		var loggerConfig zap.Config
		if strings.Compare(Env.ENVIRONMENT, "production") == 0 {
			loggerConfig = zap.NewProductionConfig()
		} else {
			loggerConfig = zap.NewDevelopmentConfig()
		}
		loggerConfig.OutputPaths = []string{"stdout", "logs/server.log"}
		loggerConfig.ErrorOutputPaths = []string{"stderr", "logs/error.log"}
		var err error
		Logger, err = loggerConfig.Build()
		if err != nil {
			log.Fatalln(err)
		}
	})
}
