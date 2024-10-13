package handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/request/keyword"
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

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
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

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.OkBooks(w, r, books)
}

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
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

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
func (h *BookHandler) LikeBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.likebook"

	log := h.bookService.Logger.With(slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	bookidstr := chi.URLParam(r, "book_id")
	bookid, err := strconv.Atoi(bookidstr)

	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}

	var book models.Book
	err = h.bookService.GetBookByID(&book, bookid)
	if err != nil {
		log.Error("book not found:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	if err := h.bookService.LikeBook(&book); err != nil {
		log.Error("failed to like book:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.StatusLikes{
		Likes:    book.Likes,
		DisLikes: book.DisLikes,
		Response: rs.OK(),
	})

}

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
func (h *BookHandler) DislikeBook(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.url.likebook"

	log := h.bookService.Logger.With(slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	bookidstr := chi.URLParam(r, "book_id")
	bookid, err := strconv.Atoi(bookidstr)

	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}

	var book models.Book
	err = h.bookService.GetBookByID(&book, bookid)
	if err != nil {
		log.Error("book not found:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	if err := h.bookService.DislikeBook(&book); err != nil {
		log.Error("failed to like book:", sl.Err(err))
		status.Err(w, r, rs.Error(err))
		return
	}

	status.Ok(w, r, rs.StatusLikes{
		Likes:    book.Likes,
		DisLikes: book.DisLikes,
		Response: rs.OK(),
	})

}

// CreateBook добавляет новую книгу
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Router /books/add [post]
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
