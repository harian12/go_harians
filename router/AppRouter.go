package router

import (
	"github.com/gin-gonic/gin"
	"go_harians/config"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func AppRouter() {
	var appConfig = AppConfig{}

	appConfig.AppName = config.GetEnv("APP_NAME", "Harians")
	appConfig.AppEnv = config.GetEnv("APP_ENV", "development")
	appConfig.AppPort = config.GetEnv("APP_PORT", "8080")

	r := gin.Default()
	RouterV1(r)
	r.Run(":" + appConfig.AppPort)
}
