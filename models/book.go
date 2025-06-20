package models

type Book struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string  `json:"title" binding:"required"`
	Author        string  `json:"author"`
	PublishedYear int     `json:"published_year"`
	Price         float64 `json:"price" binding:"required"`
}
