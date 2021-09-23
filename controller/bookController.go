package controller

import (
	"fmt"
	"net/http"
	"newApp/helper"
	"newApp/models"
	"newApp/service"
	"strconv"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(c *gin.Context)
	FindByID(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTservice
}

func NewBookController(book service.BookService,
	jwt service.JWTservice) BookController {
	return &bookController{
		bookService: book,
		jwtService:  jwt,
	}
}
func (service *bookController) All(c *gin.Context) {
	var books []models.Book
	books = service.bookService.AllBook()
	res := helper.BuildResponse(true, "ok", books)
	c.JSON(http.StatusOK, res)
}
func (service *bookController) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No Param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book models.Book = service.bookService.FindBookByID(id)
	if (book == models.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		res := helper.BuildResponse(true, "ok", book)
		c.JSON(http.StatusOK, res)
	}
}
func (service *bookController) Insert(c *gin.Context) {
	var book models.BookCreate
	err := c.ShouldBind(&book)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
		userid := service.GetUserIdByToken(authHeader)
		convertedUserId, err := strconv.ParseUint(userid, 10, 64)
		if err != nil {
			book.UserID = convertedUserId
		}
		result := service.bookService.Insert(book)
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusCreated, response)
	}

}
func (service *bookController) Update(c *gin.Context) {
	var book models.BookUpdate
	err := c.ShouldBind(&book)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
	token, err := service.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["user_id"])
	if service.bookService.IsAllowedToEdit(userid, book.ID) {
		id, err := strconv.ParseUint(userid, 10, 64)
		if err != nil {
			book.UserID = id
		}
		result := service.bookService.Update(book)
		response := helper.BuildResponse(true, "ok", result)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission ", "Not owner", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusForbidden, response)
	}

}
func (service *bookController) Delete(c *gin.Context) {
	var book models.Book
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	book.ID = id

	authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
	token, err := service.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["user_id"])
	if service.bookService.IsAllowedToEdit(userid, book.ID) {

		service.bookService.Delete(book)
		response := helper.BuildResponse(true, "ok", helper.EmptyObj{})
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission ", "Not owner", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusForbidden, response)
	}
}
func (service *bookController) GetUserIdByToken(token string) string {
	Token, err := service.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())

	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id

}
