package models

type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"firstname"`
	LastName  string `gorm:"type:varchar(255)" json:"lastname"`
	Email     string `gorm:"unique;uniqueIndex" json:"email"`
	Password  string `gorm:"size:100;->;<-;not null" json:"-"`
	Token     string `gorm:"-" json:"token,omitempty"`

	Books *[]Book `json:"books,omitempty"`
}
type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
type Register struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required"`
}
type UserUpdate struct {
	ID        uint64 `json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password,omitempty" form:"password,omitempty"`
}
