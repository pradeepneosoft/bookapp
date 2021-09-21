package repository

import (
	"newApp/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book models.Book) models.Book
	UpdateBook(book models.Book) models.Book
	DeleteBook(book models.Book)
	AllBook() []models.Book
	FindBookByID(bookID uint64) models.Book
}
type bookRepository struct {
	conn *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		conn: db,
	}
}
func (db *bookRepository) InsertBook(book models.Book) models.Book {
	db.conn.Save(&book)
	db.conn.Preload("User").Find(&book)
	return book
}
func (db *bookRepository) UpdateBook(book models.Book) models.Book {
	db.conn.Save(&book)
	db.conn.Preload("User").Find(&book)
	return book
}
func (db *bookRepository) DeleteBook(book models.Book) {
	db.conn.Delete(&book)
}
func (db *bookRepository) AllBook() []models.Book {
	var books []models.Book
	db.conn.Preload("User").Find(&books)
	return books
}
func (db *bookRepository) FindBookByID(bookID uint64) models.Book {
	var b models.Book
	db.conn.Preload("User").Find(&b, bookID)
	return b
}
