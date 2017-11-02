package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/locona/snoopy/db"
	"github.com/locona/snoopy/db/migrations"
	"github.com/locona/snoopy/middleware"
	"github.com/locona/snoopy/router"
	"github.com/spf13/viper"
)

var (
	env = flag.String("e", "dev", "環境変数")
)

func init() {
	flag.Parse()
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	switch *env {
	case "dev":
		viper.SetConfigName("dev")
	default:
		viper.SetConfigName("dev")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config error", err)
	}
}

func StartApiServer() {
	api := gin.New()
	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	api.Use(middleware.Handle400())
	api.HandleMethodNotAllowed = true

	db.SetupDB()
	db.DB.LogMode(true)
	migrations.CreateSchema()
	router.V1(api)
	router.HealthCheck(api)
	api.Run(":" + viper.GetString("port"))

}

func main() {
	StartApiServer()
}
