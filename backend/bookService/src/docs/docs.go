package docs

import (
	"bookService/src/handler"
	"github.com/go-andiamo/chioas"
	"net/http"
)

func GenDocs(commentHandler *handler.CommentHandler) *chioas.Definition {
	return &chioas.Definition{AutoHeadMethods: true,
		DocOptions: chioas.DocOptions{
			ServeDocs:       true,
			HideHeadMethods: true,
		},
		Paths: chioas.Paths{
			"/foos": {
				Methods: chioas.Methods{
					http.MethodGet: {
						Handler: commentHandler.CreateComment,
						Responses: chioas.Responses{
							http.StatusOK: {
								Description: "List of foos",
								IsArray:     true,
								SchemaRef:   "foo",
							},
						},
					},
					http.MethodPost: {
						Handler: commentHandler.CreateComment,
						Request: &chioas.Request{
							Description: "Foo to create",
							SchemaRef:   "foo",
						},
						Responses: chioas.Responses{
							http.StatusCreated: {
								Description: "New foo",
								SchemaRef:   "foo",
							},
						},
					},
					http.MethodHead: {
						Handler: commentHandler.CreateComment,
					},
				},
				Paths: chioas.Paths{
					"/{fooId}": {
						Methods: chioas.Methods{
							http.MethodGet: {
								Handler: commentHandler.CreateComment,
								Responses: chioas.Responses{
									http.StatusOK: {
										Description: "The foo",
										SchemaRef:   "foo",
									},
								},
							},
							http.MethodDelete: {
								Handler: commentHandler.CreateComment,
							},
						},
					},
				},
			},
		},
		Components: &chioas.Components{
			Schemas: chioas.Schemas{
				{
					Name:               "foo",
					RequiredProperties: []string{"name", "address"},
					Properties: chioas.Properties{
						{
							Name: "name",
							Type: "string",
						},
						{
							Name: "address",
							Type: "string",
						},
					},
				},
			},
		},
	}
}
