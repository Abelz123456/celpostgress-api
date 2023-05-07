package controller

import (
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
