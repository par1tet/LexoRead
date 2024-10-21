package handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/request/keyword"
	"bookService/src/lib/api/status"
	"bookService/src/lib/sl"
	"bookService/src/service"
	"errors"
	"fmt"
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
type limit struct {
	Limit int `json:"limit"`
}

func (s *limit) Bind(r *http.Request) error {
	if s.Limit == 0 {
		return errors.New("limit required")
	}
	return nil
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
	limitStr := chi.URLParam(r, "limit")
	limit, err := strconv.Atoi(limitStr)
	fmt.Println(limitStr)

	books, err := h.bookService.GetBooks(limit)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.Ok(w, r, books)
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

func (h *BookHandler) LikeBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.likebook"

	log := h.bookService.Logger.With(slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	bookidstr := chi.URLParam(r, "book_id")

	if err := h.bookService.LikeBook(bookidstr); err != nil {
		log.Error("failed to like book:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.OK())

}

func (h *BookHandler) DislikeBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.likebook"

	log := h.bookService.Logger.With(slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	bookid := chi.URLParam(r, "book_id")

	if err := h.bookService.DislikeBook(bookid); err != nil {
		log.Error("failed to like book:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.OK())

}

func (h *BookHandler) SearchByKeyword(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.searchByKeyword"

	log := h.bookService.Logger.With(slog.String("op", op))
	var keywordJson keyword.Keyword
	if err := render.Bind(r, &keywordJson); err != nil {
		log.Error("failed decode json", sl.Err(err))
		status.Err(w, r, rs.Error(err))
	}

	books, err := h.bookService.SearchByKeyword(keywordJson.Keyword)
	if err != nil {
		log.Error("failed to search by keyword:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
	}

	status.Ok(w, r, books)

}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.deleteBook"

	log := h.bookService.Logger.With(slog.String("op", op))
	bookId := chi.URLParam(r, "book_id")

	err := h.bookService.DeleteBook(bookId)
	if err != nil {
		log.Error("failed to delete book:", sl.Err(err))
		status.ErrWithCode(w, r, rs.Error(err), http.StatusBadRequest)
		return
	}
	status.Ok(w, r, rs.OK())
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.updateBook"
	log := h.bookService.Logger.With(slog.String("op", op))

	var book models.Book
	if err := render.Bind(r, &book); err != nil {
		log.Error("failed decode json", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	if err := h.bookService.UpdateBook(&book); err != nil {
		log.Error("failed to update book:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}
	status.Ok(w, r, rs.OK())
}
