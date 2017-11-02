package router

import (
	"github.com/gin-gonic/gin"
	csv "github.com/locona/snoopy/api/csv/v1"
	users "github.com/locona/snoopy/api/users/v1"
)

func V1(app *gin.Engine) *gin.Engine {
	v1 := app.Group("/api/v1")
	{
		csv_router := v1.Group("csv")
		csv_router.POST("upload", csv.Create)
	}
	{
		user_router := v1.Group("users")
		user_router.POST("", users.Create)
	}

	return app
}

func HealthCheck(app *gin.Engine) *gin.Engine {
	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	return app
}
