package createcomment

import (
	"bookService/src/database/initdb"
	"bookService/src/database/models"
	"bookService/src/database/repo/addcomment"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	rs.Response
	ID       int    `json:"id,omitempty"`
	BookID   int    `json:"book_id,omitempty"`
	Content  string `json:"content,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
	Likes    uint   `json:"likes,omitempty"`
	DisLikes uint   `json:"dislikes,omitempty"`
}

func New(log *slog.Logger, db *initdb.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.createcomment"

		log := log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req models.Comment
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("error on decode json", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		comment, err := addcomment.Add(&req, db)
		if err != nil {
			log.Error("failed to create comment:", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, rs.Error(err))
			return
		}

		log.Info("Comment created successfully to db")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, Response{
			Response: rs.Response{
				Status: rs.StatusOk,
				Error:  "",
			},
			ID:       comment.ID,
			BookID:   comment.BookID,
			Content:  comment.Content,
			UserID:   comment.UserID,
			DisLikes: comment.DisLikes,
			Likes:    comment.Likes,
		})
	}
}
