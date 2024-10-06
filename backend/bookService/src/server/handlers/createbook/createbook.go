package createbook

import (
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/database/repo/addbook"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	rs.Response
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Likes       uint   `json:"likes,omitempty"`    // Изменён тип на uint
	DisLikes    uint   `json:"dislikes,omitempty"` // Изменён тип на uint
	Description string `json:"description,omitempty"`
}

func New(log *slog.Logger, db *initdb.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.createbook"

		log := log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req models.Book
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("error on decode json", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		// Добавляем книгу в БД
		book, err := addbook.Add(&req, db)
		if err != nil {
			log.Error("failed to create book with comments:", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, rs.Error(err))
			return
		}

		log.Info("Book added successfully to db")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, Response{
			Response: rs.Response{
				Status: rs.StatusOk,
				Error:  "",
			},
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Likes:       book.Likes,
			DisLikes:    book.DisLikes,
			Description: book.Description,
		})
	}
}
