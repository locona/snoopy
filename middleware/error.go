package middleware

import (
	"github.com/gin-gonic/gin"
)

func Handle400() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		status := ctx.Writer.Status()
		if 400 <= status && 500 > status {
			errs := make([]error, len(ctx.Errors))
			for idx := range ctx.Errors {
				errs[idx] = ctx.Errors[idx].Err
			}
			switch status {
			case 400:
				ctx.JSON(400, gin.H{"message": "invalid request", "errors": errs})
			default:
			}
		}
	}
}
