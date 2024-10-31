package handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/status"
	"bookService/src/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RedisHandler struct {
	service *service.RedisService
}

func NewRedisHandler(s *service.RedisService) *RedisHandler {
	return &RedisHandler{service: s}
}

func (h *RedisHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "type_book")
	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid input"})
		return
	}

	if err := h.service.AddBook(&book, key); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, book)
}

// Метод для получения книги по ID
func (h *RedisHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	key := chi.URLParam(r, "type_book")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	book, err := h.service.GetBook(bookID, key)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, book)
}

func (h *RedisHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "type_book")

	books, err := h.service.GetBooks(key)

	if err != nil {
		status.Err(w, r, err)
		return
	}
	status.Ok(w, r, books)
}
func (h *RedisHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	key := chi.URLParam(r, "type_book")

	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		status.Err(w, r, err)
	}
	if err := h.service.UpdateBook(&book, key, id); err != nil {
		status.Err(w, r, err)
	}
	status.Ok(w, r, rs.OK())
}

func (h *RedisHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	key := chi.URLParam(r, "type_book")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	if err := h.service.DeleteBook(bookID, key); err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusNoContent)
}
