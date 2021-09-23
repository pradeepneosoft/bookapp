package controller

import (
	"fmt"
	"net/http"
	"newApp/helper"
	"newApp/models"
	"newApp/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Update(c *gin.Context)
	Profile(c *gin.Context)
}
type userController struct {
	UserService service.UserService
	JwtService  service.JWTservice
}

func NewUserController(userservice service.UserService, jwt service.JWTservice) UserController {
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
	authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
	token, err := u.JwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userToUpdate.ID = id
	saved := u.UserService.UpdateUser(userToUpdate)
	res := helper.BuildResponse(true, "ok", saved)
	c.JSON(http.StatusOK, res)
}
func (u *userController) Profile(c *gin.Context) {
	authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
	token, err := u.JwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := u.UserService.Profile(id)
	res := helper.BuildResponse(true, "ok", user)
	c.JSON(http.StatusOK, res)

}
