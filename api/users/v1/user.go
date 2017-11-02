package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/snoopy/db"
	"github.com/locona/snoopy/lib/errors"
	"github.com/locona/snoopy/model"
)

type signupForm struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func Create(cx *gin.Context) {
	var params signupForm
	if err := cx.BindJSON(&params); err != nil {
		cx.AbortWithError(400, err)
		return
	}

	user := model.User{
		Name:  params.Name,
		Email: params.Email,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		cx.AbortWithError(400, errors.MysqlError(err, user))
		return
	}

	cx.Writer.WriteHeader(200)
}
