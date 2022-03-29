package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN   string
	Title  string
	Price  float64
	Author string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection Error!")
	}
	db.AutoMigrate(&Book{})

	book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter"}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "PINOKIO", Price: 430.00, Author: "Carlo Collodi"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "Dear John", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	books := []Book{book1, book2, book3, book4, book5, book6}
	db.Create(&books)

	var book Book

	// Get the first record ordered by primary key and store in book variable
	// SELECT * FROM books ORDER BY id LIMIT 1;
	db.First(&book)
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)

	// Get one record, no specified order
	// SELECT * FROM books LIMIT 1;
	book = Book{}
	db.Take(&book)
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)

	// Get last record, ordered by primary key desc
	// SELECT * FROM books ORDER BY id DESC LIMIT 1;
	book = Book{}
	db.Last(&book)
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)

}
