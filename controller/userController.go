package controller

import (
	"fmt"
	"net/http"
	"newApp/helper"
	"newApp/models"
	"newApp/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(c *gin.Context)
	Profile(c *gin.Context)
}
type userController struct {
	UserService service.UserService
	JwtService  service.JwtService
}

func NewUserController(userservice service.UserService, jwt service.JwtService) UserController {
	return &userController{
		UserService: userservice,
		JwtService:  jwt,
	}
}
func (u *userController) Update(c *gin.Context) {
	var userToUpdate models.UserUpdate
	err := c.ShouldBind(&userToUpdate)
	if err != nil {
		res := helper.BuildErrorResponse("failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

}
func (u *userController) Profile(c *gin.Context) {
	authHeader := c.getHeader("Authorization")
	token, err := u.JwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "ok", user)
	c.JSON(http.StatusOK, res)

}
