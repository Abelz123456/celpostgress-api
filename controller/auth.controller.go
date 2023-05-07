package controller

import (
	"net/http"

	"github.com/celpostgress-api/common"
	"github.com/celpostgress-api/dto"
	"github.com/celpostgress-api/mapping"
	"github.com/celpostgress-api/middleware"
	"github.com/celpostgress-api/services"
	"github.com/celpostgress-api/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(authService services.IAuthService) IAuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Get All Auth
// @Summary List existing auths
// @Description Get all the existing auths
// @Tags auths
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param   authDto body  dto.PayloadLogin true "auth"
// @Router /auth/login [post]
func (controller *AuthController) Login(c *gin.Context) {
	var authDto dto.PayloadLogin
	err := c.ShouldBindJSON(&authDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}
	criteria := ""
	// criteriaQ := ""
	criteria += `where "emailName" = ` + `'` + authDto.EmailName + `'`
	// criteria += `where emailName = " + "'" + authDto.EmailName + "'"

	users := controller.AuthService.Login(c, criteria)

	err = utils.CheckPassword(authDto.Password, users.Password)
	if err != nil {
		panic(common.NewServerError(err.Error()))
	}
	marker, _ := middleware.NewJwt()
	token, _ := marker.CreateToken(users.EmailName)
	data := mapping.ToPermissionPolicyUserResponseAuth(users, token, "", "")
	response := common.ResponseOk(data, 1)
	c.JSON(http.StatusOK, response)

}

// Register
// @Summary  create user
// @Description create new user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Param   user body  dto.PayloadRegister true "user"
// @Router /auth/register [post]
func (controller *AuthController) Register(c *gin.Context) {

	var userDto dto.PayloadRegister
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}

	userResponse := controller.AuthService.Register(c, userDto)

	response := common.ResponseOk(userResponse, 1)

	c.JSON(http.StatusOK, response)
}
