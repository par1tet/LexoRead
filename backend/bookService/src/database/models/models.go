package models

type Book struct {
	ID          int    `gorm:"primaryKey"`
	Author      string `gorm:"not null" json:"author"`
	AuthorID    *uint  `gorm:"default:0" json:"author_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"size:256" json:"description"`
	Likes       uint   `gorm:"default:0"`
	DisLikes    uint   `gorm:"default:0"`
	Genre       string `gorm:"size:256" json:"genre"`

	Comments []Comment `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;" json:"comments,omitempty"`
	Files    []File    `gorm:"foreignKey:ForID;constraint:OnDelete:CASCADE;" json:"files"`
}

type File struct {
	ID         uint   `gorm:"primaryKey"`
	UploaderID *int   `gorm:"not null" json:"uploader_id"`
	FileURL    string `gorm:"not null" json:"file_url"`
	ForID      *int   `gorm:"not null" json:"for_id"`
	ForType    string `gorm:"default:book" json:"for_type"`
}

type Comment struct {
	ID       int    `gorm:"primaryKey" json:"id,omitempty"`
	BookID   int    `gorm:"not null" json:"book_id"`
	Content  string `gorm:"size:256" json:"content"`
	UserID   *int   `gorm:"default:0" json:"user_id"`
	Likes    uint64 `gorm:"default:0" json:"likes,omitempty"`
	DisLikes uint64 `gorm:"default:0" json:"dislikes,omitempty"`

	ReplyCommentID *int      `gorm:"default:null" json:"reply_comment_id,omitempty"`
	ReplyComments  []Comment `gorm:"foreignKey:ReplyCommentID;references:ID" json:"reply_comments,omitempty"`
}
