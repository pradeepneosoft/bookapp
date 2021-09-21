package service

import (
	"fmt"
	"log"
	"newApp/models"
	"newApp/repository"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(book models.BookCreate) models.Book
	Update(book models.BookUpdate) models.Book
	Delete(book models.Book)
	AllBook() []models.Book
	FindBookByID(BookID uint64) models.Book
	IsAllowedToEdit(UserID string, BookID uint64) bool
}
type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: repo,
	}
}

func (service *bookService) Insert(book models.BookCreate) models.Book {
	b := models.Book{}
	err := smapping.FillStruct(&b, smapping.MapFields(book))
	if err != nil {
		log.Fatalf("failed to map %v:", err)
	}
	res := service.bookRepository.InsertBook(b)
	return res

}
func (service *bookService) Update(book models.BookUpdate) models.Book {
	b := models.Book{}
	err := smapping.FillStruct(&b, smapping.MapFields(book))
	if err != nil {
		log.Fatalf("failed to map %v:", err)
	}
	res := service.bookRepository.UpdateBook(b)
	return res

}
func (service *bookService) Delete(book models.Book) {
	service.bookRepository.DeleteBook(book)
}
func (service *bookService) AllBook() []models.Book {
	return service.bookRepository.AllBook()
}
func (service *bookService) FindBookByID(BookID uint64) models.Book {
	return service.bookRepository.FindBookByID(BookID)
}
func (service *bookService) IsAllowedToEdit(UserID string, BookID uint64) bool {
	b := service.bookRepository.FindBookByID(BookID)
	id := fmt.Sprintf("%v", b.UserID)
	return id == UserID
}
