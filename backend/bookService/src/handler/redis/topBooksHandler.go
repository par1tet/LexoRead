package redis_handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/status"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *RedisHandler) GetTopBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetTopBooks()
	if err != nil {
		status.Err(w, r, err)
		return
	}

	render.JSON(w, r, books)
}

func (h *RedisHandler) AddTopBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		status.Err(w, r, err)
		return
	}

	if err := h.service.AddToTopBooks(&book); err != nil {
		status.Err(w, r, rs.Error(err))

		return
	}

	status.Ok(w, r, book)
}

func (h *RedisHandler) DeleteTopBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}

	if err := h.service.DeleteTopBook(bookID); err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.OK())
}
