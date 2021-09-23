package middleware

import (
	"log"
	"net/http"
	"newApp/helper"
	"newApp/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJwt(jwtService service.JWTservice) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := helper.GetTokenFromHeader(c.GetHeader("Authorization"))
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("cliams[user_id]", claims["user_id"])
			log.Println("cliams[issuer]", claims)

		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token Is Not Valid", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
