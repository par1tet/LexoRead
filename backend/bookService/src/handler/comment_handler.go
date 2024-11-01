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

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := render.Bind(r, &comment); err != nil {
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
	bookIDStr := chi.URLParam(r, "book_id")
	bookId, err := strconv.Atoi(bookIDStr)
	if err != nil {
		status.Err(w, r, rs.Error(err))
	}

	comments, err := h.commentService.GetCommentsByBookID(bookId)
	if err != nil {
		status.Err(w, r, rs.Error(err))
	}

	status.Ok(w, r, comments)
}

func (h *CommentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	comment, err := h.commentService.GetCommentByID(id)
	if err != nil {
		status.Err(w, r, rs.Error(err))
		return
	}
	status.Ok(w, r, comment)

}
