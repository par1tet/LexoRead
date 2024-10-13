package models

import (
	"errors"
	"net/http"
)

func (b *Book) Bind(r *http.Request) error {
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.AuthorID == nil {
		return errors.New("author_id is required")
	}
	if b.Files == nil {
		return errors.New("files is required")
	}

	return nil
}

func (b *File) Bind(r *http.Request) error {
	if b.UploaderID == nil {
		return errors.New("uploader_id is required")
	}
	if b.FileURL == "" {
		return errors.New("file_url is required")
	}
	if b.ForID == nil {
		return errors.New("for_id is required")
	}

	return nil
}

func (b *Comment) Bind(r *http.Request) error {
	if b.Content == "" {
		return errors.New("content is required")
	}

	if b.UserID == nil {
		return errors.New("for_id is required")
	}

	return nil
}
