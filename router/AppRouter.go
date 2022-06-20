package router

import (
	"github.com/gin-gonic/gin"
	"go_harians/app/helpers"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func AppRouter() {
	var appConfig = AppConfig{}

	appConfig.AppName = helpers.GetEnv("APP_NAME", "Harians")
	appConfig.AppEnv = helpers.GetEnv("APP_ENV", "development")
	appConfig.AppPort = helpers.GetEnv("APP_PORT", "8080")

	r := gin.Default()
	RouterV1(r)
	r.Run(":" + appConfig.AppPort)
}
