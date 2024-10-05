package models

type Book struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Author string
	Likes  uint `gorm:"default:0"`
}
