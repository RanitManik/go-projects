package models

import (
	"github.com/RanitManik/go-projects/03-book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Publication string `gorm:"type:varchar(255)" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() error {
	result := db.Create(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	result := db.Find(&Books)
	if result.Error != nil {
		return nil, result.Error
	}

	return Books, nil
}

func GetBookById(id int64) (*Book, error) {
	var book Book

	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func DeleteBookById(id int64) error {
	result := db.Delete(&Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateBookById(id int64, updated *Book) (*Book, error) {
	var book Book

	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if updated.Name != "" {
		book.Name = updated.Name
	}
	if updated.Author != "" {
		book.Author = updated.Author
	}
	if updated.Publication != "" {
		book.Publication = updated.Publication
	}

	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
