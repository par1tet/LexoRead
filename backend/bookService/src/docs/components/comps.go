package components

import (
	"bookService/src/handler"
	redis_handler "bookService/src/handler/redis"
	"net/http"

	chioas "github.com/go-andiamo/chioas"
)

comps := chioas.Components
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
