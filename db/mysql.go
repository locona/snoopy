package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

var (
	HOST         string
	PORT         string
	USER         string
	PASSWORD     string
	DATABASE     string
	DATABASE_URI string
)

// MySQLのセットアップ
func SetupDB() {
	if DB != nil {
		return
	}

	HOST = viper.GetString("mysql.host")
	PORT = viper.GetString("mysql.port")
	USER = viper.GetString("mysql.user")
	PASSWORD = viper.GetString("mysql.password")
	DATABASE = viper.GetString("mysql.database")
	DATABASE_URI = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		USER, PASSWORD, HOST, PORT, DATABASE,
	)
	mysql, err := gorm.Open("mysql", DATABASE_URI)
	if err != nil {
		fmt.Println("MYSQL ERROR:", err)
		os.Exit(1)
	}

	DB = mysql
}
