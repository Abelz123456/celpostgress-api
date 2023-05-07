package common

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err interface{}) {

	if NotFoundErrors(c, err) {
		return
	}

	if ValidationErrors(c, err) {
		return
	}

	InternalServerError(c, err)
}

func ValidationErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)
	log.Println(err)

	if ok {
		response := ResponseOkDataNotFound(exception.Error)
		c.JSON(http.StatusOK, response)
		return true
	} else {
		return false
	}
}

func NotFoundErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		response := ResponseOkDataNotFound(exception.Error)
		c.JSON(http.StatusOK, response)
		return true
	} else {
		return false
	}
}

func InternalServerError(c *gin.Context, err interface{}) {
	response := ResponseInternalServerError(err)
	c.JSON(http.StatusOK, response)
}
