package docs

import "newApp/models"

// swagger:route GET /books books-tag
// Books return list of books.
// responses:
//   200: {[]models.Book}

// This text will appear as description of your response body.
// swagger:response bookResponse
type booksResponseWrapper struct {
	// in:body
	Body []models.Book
}

type bookParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	// Body api.FooBarRequest
}
