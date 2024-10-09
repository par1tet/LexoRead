// src/handler/comment_handler.go
package handler

import (
	"bookService/src/database/models"
	rs "bookService/src/lib/api/request"
	"bookService/src/lib/api/status"
	"bookService/src/service"
	"net/http"

	"github.com/go-chi/render"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := render.DecodeJSON(r.Body, &comment); err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	if err := h.commentService.CreateComment(&comment); err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.Ok(w, r, rs.OK())
}

func (h *CommentHandler) GetCommentsByBookID(w http.ResponseWriter, r *http.Request) {
	// Реализуйте получение комментариев по ID книги
}
