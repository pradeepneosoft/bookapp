package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTservice interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
type JwtService struct {
	SecretKey string
	Issuer    string
}

func NewJWTservice() JWTservice {
	return &JwtService{
		SecretKey: getSecretKey(),
		Issuer:    "newApp",
	}
}
func getSecretKey() string {
	secretkey := os.Getenv("JWT_SECRET")
	if secretkey == "" {
		secretkey = "newApp"
	}
	return secretkey

}

func (jwtservice *JwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwtservice.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtservice.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (jwtservice *JwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t.Method.(*jwt.SigningMethodHMAC))
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println(t.Header)
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(jwtservice.SecretKey), nil
	})
}
