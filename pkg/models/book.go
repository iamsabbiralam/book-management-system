package models

import (
	"practice/book-management-system/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var GetBook Book
	db := db.Where("Id = ?", Id).First(&GetBook)
	return &GetBook, db
}

func DeleteBook(ID int64) Book {
	var Book Book
	db.Where("ID = ?", ID).Delete(&Book)
	return Book
}
