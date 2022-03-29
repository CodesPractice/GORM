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
		panic("Database conncetion failure!!")
	}

	db.AutoMigrate(&Book{}) // make db schema from the struct

	book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter"}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "PINOKIO", Price: 430.00, Author: "Carlo Collodi"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "Dear John", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	//Batch Insert
	books := []Book{book1, book2, book3, book4, book5, book6}
	db.Create(books)

	//Create From Map
	//GORM supports create from map[string]interface{} and []map[string]interface{}{},
	db.Model(&Book{}).Create(map[string]interface{}{
		"ISBN": "1007-KIDS-777", "Title": "Harry Potter III", "Price": 1430.00, "Author": "J. K. Rowling",
	})

	var book Book
	db.First(&book) // get the first inserted record from the db table
	fmt.Println(book.ISBN, book.Title, book.Price, book.Author)

	book = Book{}  // empty the pointer of book
	db.Last(&book) // get the last inserted record from the db table
	fmt.Println(book.ISBN, book.Title, book.Price, book.Author)

}
