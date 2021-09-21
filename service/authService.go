package service

import (
	"log"
	"newApp/helper"
	"newApp/models"
	"newApp/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user models.Register) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}
type authService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.UserRepo.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		Check := helper.CheckPasswordHash(v.Password, password)
		if Check && v.Email == email {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user models.Register) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed to map %v", err)
	}
	res := service.UserRepo.InsertUser(userToCreate)
	return res
}
func (service *authService) FindByEmail(email string) models.User {
	return service.UserRepo.FindByEmail(email)

}
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.UserRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
