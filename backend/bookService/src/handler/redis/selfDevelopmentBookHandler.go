package redis_handler

import (
	"bookService/src/database/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *RedisHandler) AddSelfDevelopmentBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid input"})
		return
	}

	if err := h.service.AddSelfDevelopmentBook(&book); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, book)
}

func (h *RedisHandler) GetSelfDevelopmentBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetSelfDevelopmentBooks()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, books)
}

func (h *RedisHandler) GetSelfDevelopmentBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	book, err := h.service.GetSelfDevelopmentBook(bookID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, book)
}

func (h *RedisHandler) DeleteSelfDevelopmentBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	if err := h.service.DeleteSelfDevelopmentBook(bookID); err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusNoContent)
}
