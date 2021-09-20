package controller

import (
	"net/http"
	"newApp/helper"
	"newApp/models"
	"newApp/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(c *gin.Context)
}
type userController struct {
	UserService service.UserService
}

func NewUserController(userservice service.UserService) UserController {
	return &userController{
		UserService: userservice,
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
