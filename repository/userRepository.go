package repository

import (
	"newApp/helper"
	"newApp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	Updateuser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) *gorm.DB
	FindByEmail(email string) models.User
	// ProfileUser(userId string) models.User
}

type UserConnection struct {
	Conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnection{
		Conn: db,
	}
}
func (db *UserConnection) InsertUser(user models.User) models.User {
	user.Password = helper.HashPassword([]byte(user.Password))
	db.Conn.Save(&user)
	return user
}
func (db *UserConnection) Updateuser(user models.User) models.User {
	if user.Password != "" {
		user.Password = helper.HashPassword([]byte(user.Password))
	} else {
		var u models.User
		db.Conn.Find(&u, user.ID)
		user.Password = u.Password
	}
	db.Conn.Save(&user)
	return user
}
func (db *UserConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.Conn.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (db *UserConnection) IsDuplicateEmail(email string) *gorm.DB {
	var user models.User
	return db.Conn.Where("email = ?", email).Take(&user)
}
func (db *UserConnection) FindByEmail(email string) models.User {
	var user models.User
	db.Conn.Where("email = ?", email).Take(&user)
	return user
}

// func (db *UserConnection) ProfileuUser(userId string) models.User {

// }
