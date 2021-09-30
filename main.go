package main

import (
	"fmt"
	"newApp/config"
	"newApp/controller"
	"newApp/middleware"
	"newApp/repository"
	"newApp/service"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	//_ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/docs"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTservice        = service.NewJWTservice()
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	Router                                   = gin.Default()
)

// @title newApp Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email pradeep.bora@neosoftmail.com

// @BasePath /api/v1
func main() {
	fmt.Println("sever up & started")
	defer config.ClosedatabaseConnection(db)
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := Router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := Router.Group("api/user", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.PUT("profile", userController.Update)
		userRoutes.GET("profile", userController.Profile)

	}
	bookRoutes := Router.Group("api/books", middleware.AuthorizeJwt(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	Router.Run(":8000")
}
