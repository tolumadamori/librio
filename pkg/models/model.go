package models

import (
	"gorm.io/gorm"
)

// This is the model all the books are built on
type Book struct {
	ID     int    `gorm:"not null" json:"id,string"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Type   string `json:"type"`
}

// Create Book model. Creates a new book in the DB.
func (b *Book) CreateBook(db *gorm.DB) *Book {

	db.Create(b)
	return b

}

// find book model. Returns all books
func FindBooks(db *gorm.DB) []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Find book model. Finds a specific book by Id.
func FindBookByID(db *gorm.DB, ID int64) *Book {
	var BookWithID Book
	db.Where("ID=?", ID).Find(&BookWithID)
	return &BookWithID
}

// Update book model. Updates the book details for the book with the Id passed in.
func UpdateBook(db *gorm.DB, ID int64, Booktoupdate Book) *Book {
	title, author, genre := Booktoupdate.Title, Booktoupdate.Author, Booktoupdate.Type
	db.Where("ID=?", ID).Find(&Booktoupdate).Model(&Booktoupdate).Updates(map[string]interface{}{"ID": ID, "title": title, "author": author, "type": genre})
	return &Booktoupdate
}

// Delete book model. Deletes a particular book from the DB.
func DeleteBook(db *gorm.DB, ID int64) *Book {
	var BookWithID Book
	db.Where("ID=?", ID).Find(&BookWithID).Delete(&BookWithID)
	return &BookWithID
}
