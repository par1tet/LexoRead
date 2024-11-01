package docs

import (
	"bookService/src/handler"
	"net/http"

	"github.com/go-andiamo/chioas"
)

func GenDocs(bookHandler *handler.BookHandler, redisHandler *handler.RedisHandler, commentHandler *handler.CommentHandler) *chioas.Definition {
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
			"/books/search/{query}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler:     bookHandler.SearchByKeyword,
						Description: "Search book",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book deleted",
								IsArray:     true,
								SchemaRef:   "Book",
							},
							http.StatusNoContent: {
								Description: "Ничего не найдено",
								SchemaRef:   "error",
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
							SchemaRef:   "BookWithoutId",
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
			"/books/{book_id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodDelete: {
						Handler:     bookHandler.DeleteBook,
						Description: "Delete book",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book deleted",
								SchemaRef:   "status",
							},
							http.StatusBadRequest: {
								Description: "Bad Request",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/books": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPut: {
						Handler:     bookHandler.UpdateBook,
						Description: "Update book",
						Request: &chioas.Request{
							Description: "Book to update",
							SchemaRef:   "Book",
						},
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book deleted",
								SchemaRef:   "status",
							},
							http.StatusBadRequest: {
								Description: "Bad Request",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/books/like/{book_id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost: {
						Handler:     bookHandler.LikeBook,
						Description: "Like book",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book like",
								SchemaRef:   "status",
							},
							http.StatusBadRequest: {
								Description: "Bad Request",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/books/dislike/{book_id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost: {
						Handler:     bookHandler.DislikeBook,
						Description: "Dislike book",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Book dislike",
								SchemaRef:   "status",
							},
							http.StatusBadRequest: {
								Description: "Bad Request",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/redis/{type_book}/{id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler: redisHandler.GetBook,
						Description: "get book by redis type and id redis_types:top\n" +
							"default\nbusiness\nfiction\nliterature\nnovelty\npopular\npsychology\n" +
							"romance\nself-development",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Ok",
								SchemaRef:   "status",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
					http.MethodDelete: {
						Handler: redisHandler.DeleteBook,
						Description: "delete book by redis type and id redis_types:top\n" +
							"default\nbusiness\nfiction\nliterature\nnovelty\npopular\npsychology\n" +
							"romance\nself-development",
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Ok",
								SchemaRef:   "status",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
					http.MethodPut: {
						Handler: redisHandler.UpdateBook,
						Description: "update book by redis type and id redis_types:top\n" +
							"default\nbusiness\nfiction\nliterature\nnovelty\npopular\npsychology\n" +
							"romance\nself-development",
						Request: &chioas.Request{
							Description: "Book to update",
							SchemaRef:   "BookWithoutId",
						},
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Ok",
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
			"/redis/{type_book}/": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler: redisHandler.GetBooks,
						Description: "get books by redis_types:top\n" +
							"default\nbusiness\nfiction\nliterature\nnovelty\npopular\npsychology\n" +
							"romance\nself-development",
						Responses: chioas.Responses{
							http.StatusOK: {
								IsArray:     true,
								Description: "Ok",
								SchemaRef:   "Book",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
			"/redis/{type_book}/add": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost: {
						Handler: redisHandler.AddBook,
						Description: "add book by redis type redis_types:top\n" +
							"default\nbusiness\nfiction\nliterature\nnovelty\npopular\npsychology\n" +
							"romance\nself-development",
						Request: &chioas.Request{
							Description: "Book to add",
							SchemaRef:   "BookWithoutId",
						},
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Ok",
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
			"/comments/add": chioas.Path{
				Methods: chioas.Methods{
					http.MethodPost: {
						Handler:     commentHandler.CreateComment,
						Description: "add comment to database",
						Request: &chioas.Request{
							Description: "Comment to create",
							SchemaRef:   "Comment",
						},
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "Ok",
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
			"/comments/book_id={book_id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler:     commentHandler.GetCommentsByBookID,
						Description: "get comments by book id",
						Responses: chioas.Responses{
							http.StatusOK: {
								IsArray:     true,
								Description: "list of comments",
								SchemaRef:   "Comment",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
				},
			},

			"/comments/{id}": chioas.Path{
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler:     commentHandler.GetCommentsByBookID,
						Description: "get comments by id",
						Responses: chioas.Responses{
							http.StatusOK: {
								IsArray:     true,
								Description: "list of comments",
								 SchemaRef:   "Comment",
							},
							http.StatusInternalServerError: {
								Description: "Error",
								SchemaRef:   "error",
							},
						},
					},
				},
			},
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
					Name: "BookWithoutId",
					Properties: chioas.Properties{
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
						{
							Name: "reply_comment_id",
							Type: "integer",
						},
						{
							Name:      "reply_comments",
							Type:      "array",
							SchemaRef: "Comment_for_reply_comments",
						},
					},
				},
				{
					Name: "Comment_for_reply_comments",
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
						{
							Name: "reply_comment_id",
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
