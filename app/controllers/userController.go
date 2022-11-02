package controllers

import (
	"github.com/TendonT52/go-tendon-api/app/models"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
	models.SignUp(ctx)
	if len(ctx.Errors) > 0 {
		return 
	}
}
