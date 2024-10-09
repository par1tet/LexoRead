package models

import (
	"errors"
	"net/http"
)

type Book struct {
	ID          int    `gorm:"primaryKey" json:"id,omitempty"`
	Author      string `gorm:"not null" json:"author" binding:"require"`
	AuthorID    *uint  `gorm:"default:0" json:"author_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"size:256" json:"description,omitempty"`
	Likes       uint   `gorm:"default:0" json:"likes,omitempty"`
	DisLikes    uint   `gorm:"default:0" json:"dislikes,omitempty"`
	BookFileID  uint   `gorm:"default:0" json:"book_file_id,omitempty"`

	Comments []Comment `gorm:"foreignKey:BookID" json:"comments,omitempty"`
}

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"size:15"`
	Real_name string
	AvatarURL string `gorm:"default:null"`
}

type File struct {
	ID         uint   `gorm:"primaryKey" json:"id,omitempty"`
	UploaderID string `gorm:"not null" json:"uploader_id"`
	FileURL    string `gorm:"not null" json:"file_url,omitempty"`
	ForID      int    `gorm:"not null" json:"for_id"`
	ForType    string `gorm:"default:book" json:"for_type"`
}

type Comment struct {
	ID       int    `gorm:"primaryKey" json:"id,omitempty"`
	BookID   int    `gorm:"not null" json:"book_id"`
	Content  string `gorm:"size:256" json:"content"`
	UserID   int    `gorm:"default:0" json:"user_id"`
	Likes    uint64 `gorm:"default:0" json:"likes,omitempty"`
	DisLikes uint64 `gorm:"default:0" json:"dislikes,omitempty"`

	ReplyCommentID *int      `gorm:"default:null" json:"reply_comment_id,omitempty"`
	ReplyComments  []Comment `gorm:"foreignKey:ReplyCommentID;references:ID" json:"reply_comments,omitempty"`
}

func (b *Book) Bind(r *http.Request) error {
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.AuthorID == nil {
		return errors.New("title is required")
	}
	return nil
}
