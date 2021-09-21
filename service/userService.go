package service

import (
	"log"
	"newApp/models"
	"newApp/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	UpdateUser(user models.UserUpdate) models.User
	Profile(userID string) models.User
}
type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}

}
func (service *userService) UpdateUser(user models.UserUpdate) models.User {
	userToUpadte := models.User{}
	err := smapping.FillStruct(&userToUpadte, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	updatedUser := service.UserRepo.Updateuser(userToUpadte)
	return updatedUser
}
func (service *userService) Profile(userID string) models.User {
	return service.UserRepo.ProfileUser(userID)
}
