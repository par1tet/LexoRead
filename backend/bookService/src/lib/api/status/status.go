package status

import (
	"bookService/src/database/models"
	"net/http"

	"github.com/go-chi/render"
)

func OkBooks(w http.ResponseWriter, r *http.Request, json []models.Book) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, json)
}

func Ok(w http.ResponseWriter, r *http.Request, json interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, json)
}

func Err(w http.ResponseWriter, r *http.Request, err interface{}) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, err)
}
