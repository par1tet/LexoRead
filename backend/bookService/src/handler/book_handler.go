package handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/status"
	"bookService/src/lib/sl"
	"bookService/src/service"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.createbook"

	log := h.bookService.Logger.With(slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	var req models.Book

	if err := render.Bind(r, &req); err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}

	log.Info("request body decoded", slog.Any("request", req))

	if err := h.bookService.CreateBook(&req); err != nil {
		log.Error("failed to create book with comments:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.OK())
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.OkBooks(w, r, books)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	bookIDStr := chi.URLParam(r, "book_id")
	h.bookService.Logger.Info("bookIDStr", slog.String("bookIDStr", bookIDStr))
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	err = h.bookService.GetBookByID(&book, bookID)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.Ok(w, r, book)
}

// func GetBookByID(w http)
