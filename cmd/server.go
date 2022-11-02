package main

import (
	"net/http"

	"github.com/TendonT52/go-tendon-api/app/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "ok", Message: "pong"})
	})

	router.POST("/sign-in", controllers.SignUpHandler)

	router.Run(":8080")
}
