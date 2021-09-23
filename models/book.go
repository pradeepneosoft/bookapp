package models

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
type BookUpdate struct {
	ID          uint64 `form:"id" json:"id" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
type BookCreate struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
