package controller

import (
	"net/http"
	"newApp/helper"
	"newApp/models"
	"newApp/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
type authController struct {
	authService service.AuthService
	jwtService  service.JWTservice
}

func NewAuthController(auth service.AuthService, jwt service.JWTservice) AuthController {
	return &authController{
		authService: auth,
		jwtService:  jwt,
	}
}

func (auth *authController) Login(c *gin.Context) {
	var login models.Login
	err := c.ShouldBind(&login)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := auth.authService.VerifyCredential(login.Email, login.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := auth.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "ok", v)
		c.JSON(http.StatusOK, response)

	}
	response := helper.BuildErrorResponse("Invalid credential ", "Invelid credential", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
func (auth *authController) Register(c *gin.Context) {
	var register models.Register
	err := c.ShouldBind(&register)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !auth.authService.IsDuplicateEmail(register.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		createduser := auth.authService.CreateUser(register)
		token := auth.jwtService.GenerateToken(strconv.FormatUint(createduser.ID, 10))
		createduser.Token = token
		response := helper.BuildResponse(true, "ok", createduser)
		c.JSON(http.StatusCreated, response)

	}
}
