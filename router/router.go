package router

import (
	"net/http"
	"x-men/src"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Test Meli",
		})
	})

	r.POST("/mutant", func(context *gin.Context) {
		src.Mutant(context)
	})
}
