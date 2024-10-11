package redis_models

type Book struct {
	ID          int    `json:"id"` // ID - строка
	Title       string `json:"title"`
	Author      string `json:"author"`
	AuthorID    int    `json:"author_id"`
	Description string `json:"description"`
	Files       []File `json:"files"`
}

type File struct {
	UploaderID int    `json:"uploader_id"`
	FileURL    string `json:"file_url"`
	ForID      int    `json:"for_id"`
	ForType    string `json:"for_type"`
}
