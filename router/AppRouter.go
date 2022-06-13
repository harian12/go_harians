package router

import (
	"github.com/gin-gonic/gin"
)

func AppRouter() {
	r := gin.Default()
	RouterV1(r)
	r.Run()
}
