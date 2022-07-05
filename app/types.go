package app

import "time"

type Book struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title       string `gorm:"type:varchar(30);" json:"title"`
	Description string `gorm:"type:text;" json:"description"`
	Price       int    `gorm:"type:int;" json:"price"`
	Rating      int    `gorm:"type:int;" json:"rating"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
