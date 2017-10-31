package router

import (
	"github.com/gin-gonic/gin"
	csv "github.com/locona/snoopy/api/csv/v1"
)

func V1(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/api/v1")
	{
		csv_router := v1.Group("csv")
		csv_router.POST("upload", csv.Create)
	}

	return app
}

func HealthCheck(app *gin.Engine) *gin.Engine {
	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok", "web-token": webToken})
	})

	return app
}
