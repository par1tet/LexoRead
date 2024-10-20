package docs

import (
	"bookService/src/handler"
	"net/http"

	"github.com/go-andiamo/chioas"
)

func GenDocs(bookHandler *handler.BookHandler) *chioas.Definition {
	return &chioas.Definition{
		DocOptions: chioas.DocOptions{
			UIStyle:        chioas.Swagger,
			Title:          "Book",
			SwaggerOptions: chioas.SwaggerOptions{},
			SpecName:       "bookService",
			ServeDocs:      true,
		},
		Paths: chioas.Paths{
			"/books/id/{book_id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler:     bookHandler.GetBookByID,
						Description: "Return one book by unique id",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book",
								SchemaRef:   "Book",
							},
						},
					},
				},
			},
			"/books/limit={limit}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler:     bookHandler.GetBooks,
						Description: "Returns a list of books with an optional limit",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "List of books",
								IsArray:     true,
								SchemaRef:   "Book",
							},
						},
					},
				},
			},
			"/books/add": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost: {
						Handler:     bookHandler.GetBooks,
						Description: "Add book to database",
						Request: &chioas.Request{
							Description: "Book to create",
							SchemaRef:   "Book",
						},
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "return status",
								SchemaRef:   "status",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/books/delete": chioas.Path{},
		},
		Components: &chioas.Components{
			Schemas: chioas.Schemas{
				{
					Name: "error",
					Properties: chioas.Properties{
						{
							Name: "status",
							Type: "string",
						},
						{
							Name: "error",
							Type: "string",
						},
					},
				},
				{
					Name: "status",
					Properties: chioas.Properties{
						{
							Name: "status",
							Type: "string",
						},
					},
				},
				{
					Name: "Limit",
					Properties: chioas.Properties{
						{
							Name: "limit",
							Type: "integer",
						},
					},
				},
				{
					Name: "Book",
					Properties: chioas.Properties{
						{
							Name: "id",
							Type: "integer",
						},
						{
							Name: "author",
							Type: "string",
						},
						{
							Name: "author_id",
							Type: "integer",
						},
						{
							Name: "title",
							Type: "string",
						},
						{
							Name: "description",
							Type: "string",
						},
						{
							Name: "likes",
							Type: "integer",
						},
						{
							Name: "dislikes",
							Type: "integer",
						},
						{
							Name:      "comments",
							Type:      "array",
							SchemaRef: "Comment", // Ссылка на схему "Comment"
						},
						{
							Name:      "files",
							Type:      "array",
							SchemaRef: "File", // Ссылка на схему "File"
						},
					},
				},
				{
					Name: "Comment",
					Properties: chioas.Properties{
						{
							Name: "id",
							Type: "integer",
						},
						{
							Name: "book_id",
							Type: "integer",
						},
						{
							Name: "content",
							Type: "string",
						},
						{
							Name: "user_id",
							Type: "integer",
						},
					},
				},
				{
					Name: "File",
					Properties: chioas.Properties{
						{
							Name: "id",
							Type: "integer",
						},
						{
							Name: "uploader_id",
							Type: "integer",
						},
						{
							Name: "file_url",
							Type: "string",
						},
						{
							Name: "for_id",
							Type: "integer",
						},
						{
							Name: "for_type",
							Type: "string",
						},
					},
				},
			},
		},
	}
}
