package slice

import "bookService/src/database/models"

func RemoveFirstBook(s []models.Book) []models.Book {
	if len(s) == 0 {
		return s
	}
	return s[1:]
}
