package main

import (
	"fmt"
	"newApp/config"
	"newApp/controller"
	"newApp/middleware"
	"newApp/repository"
	"newApp/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
)

func main() {
	fmt.Println("sever up & started")
	defer config.ClosedatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/user", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.PUT("profile", userController.Update)
		userRoutes.GET("profile", userController.Profile)

	}
	bookRoutes := r.Group("api/books", middleware.AuthorizeJwt(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	r.Run(":8000")
}
