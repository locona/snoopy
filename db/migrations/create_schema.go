package migrations

import (
	"github.com/locona/snoopy/db"
	"github.com/locona/snoopy/model"
)

func CreateSchema() {
	db.DB.AutoMigrate(
		&model.User{},
	)
}
