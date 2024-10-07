package addcomment

import (
	"bookService/src/database/initdb"
	"bookService/src/database/models"
)

func Add(comment *models.Comment, db *initdb.DB) (*models.Comment, error) {
	err := db.DB.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil

}
