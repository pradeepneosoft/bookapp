package service_test

import (
	// . "practice/newApp/service"

	"newApp/config"
	"newApp/models"

	// . "newApp/models"
	"newApp/repository"
	"newApp/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var book models.BookCreate = models.BookCreate{
	// ID:          10,
	Title:       "something",
	Description: "Something more",
	UserID:      1,
}
var _ = Describe("BookService", func() {
	var (
		db             *gorm.DB
		bookRepository repository.BookRepository
		bookService    service.BookService
	)
	BeforeSuite(func() {
		db = config.SetupDatabaseConnection()
		bookRepository = repository.NewBookRepository(db)
		bookService = service.NewBookService(bookRepository)
	})
	Describe("fetching one book", func() {

		Context("if there is atleast one video", func() {
			BeforeEach(func() {
				bookService.Insert(book)
			})
			It("it should return atleast one record", func() {
				list := bookService.FindBookByID(4)
				Expect(list).ShouldNot(BeEmpty())
			})
			It("should map feild correctly", func() {
				firstRecord := bookService.AllBook()[0]
				Expect(firstRecord.Title).Should(Equal(book.Title))
				Expect(firstRecord.Description).Should(Equal(book.Description))
				Expect(firstRecord.UserID).Should(Equal(book.UserID))

			})
		})
		Context("if there is no such id ", func() {
			It("Should return empty list ", func() {
				emptyList := bookService.FindBookByID(5000)
				Expect(emptyList).Should(BeEmpty())
			})
		})
	})
	Describe("frtching all books", func() {

		Context("if there is atleast one video", func() {
			BeforeEach(func() {
				bookService.Insert(book)
			})
			It("it should return atleast one record", func() {
				list := bookService.AllBook()
				Expect(list).ShouldNot(BeEmpty())
			})
			It("should map feild correctly", func() {
				firstRecord := bookService.AllBook()[0]
				Expect(firstRecord.Title).Should(Equal(book.Title))
				Expect(firstRecord.Description).Should(Equal(book.Description))
				Expect(firstRecord.UserID).Should(Equal(book.UserID))

			})
		})
		Context("if there is no video", func() {
			It("Should return empty list ", func() {
				emptyList := bookService.AllBook()
				Expect(emptyList).Should(BeEmpty())
			})
		})
	})

	AfterSuite(func() {
		config.ClosedatabaseConnection(db)
	})
})
