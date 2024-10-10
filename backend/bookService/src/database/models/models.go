package models

type Book struct {
	ID          int    `gorm:"primaryKey" json:"id,omitempty"`
	Author      string `gorm:"not null" json:"author"`
	AuthorID    uint   `gorm:"default:0" json:"author_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"size:256" json:"description,omitempty"`
	Likes       uint   `gorm:"default:0" json:"likes,omitempty"`
	DisLikes    uint   `gorm:"default:0" json:"dislikes,omitempty"`
	BookFileID  uint   `gorm:"default:0" json:"book_file_id,omitempty"`

	Comments []Comment `gorm:"foreignKey:BookID" json:"comments,omitempty"` // Добавили тег JSON
}

type File struct {
	ID         int    `gorm:"primaryKey"`
	UploaderID uint   `gorm:"default:0"`
	FileURL    string `gorm:"not null"`
	ForID      int    `gorm:"not null"`
	For        string `gorm:"default:book"`
}

type Comment struct {
	ID       int    `gorm:"primaryKey"`
	BookID   int    `gorm:"not null"`
	Content  string `gorm:"size:256"`
	UserID   int    `gorm:"default:0"`
	Likes    uint   `gorm:"default:0"`
	DisLikes uint   `gorm:"default:0"`
}
