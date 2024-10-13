package handler

import (
	"bookService/src/database/models"
	"bookService/src/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Определение структуры RedisHandler
type RedisHandler struct {
	service *service.RedisService // Используем RedisService
}

// Конструктор для RedisHandler
func NewRedisHandler(s *service.RedisService) *RedisHandler {
	return &RedisHandler{service: s}
}

// Метод для добавления книги
func (h *RedisHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid input"})
		return
	}

	if err := h.service.AddBook(&book); err != nil {
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
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	book, err := h.service.GetBook(bookID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, book)
}

// Метод для получения топ-100 книг
func (h *RedisHandler) GetTopBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetTopBooks()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, books)
}

// Метод для удаления книги
func (h *RedisHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid book ID"})
		return
	}

	if err := h.service.DeleteBook(bookID); err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusNoContent)
}
