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

	var book Book
	var books []Book

	book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter"}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "PINOKIO", Price: 430.00, Author: "Carlo Collodi"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "Dear John", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	books = []Book{book1, book2, book3, book4, book5, book6}
	db.Create(&books)

	// SELECT * FROM books WHERE id = 3;
	db.First(&book, 3)
	fmt.Println("----- Get book of id 3 -----")
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)
	/* 	Title           : PINOKIO
	   	ISBN            : 1003-KIDS-333
	   	Author          : Carlo Collodi
	   	Price           : USD 430.00   */

	// SELECT * FROM books WHERE id = 1;
	book = Book{} // clean the variable
	db.First(&book, "1")
	fmt.Println("----- Get book of id 1 -----")
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)
	/* 	Title           : PINOKIO
	   	ISBN            : 1003-KIDS-333
	   	Author          : Carlo Collodi
	   	Price           : USD 430.00  */

	// SELECT * FROM book WHERE id IN (1,2,3);
	db.Find(&books, []int{2, 4})
	fmt.Println("----- Get the books of id 2 and 4 -----")
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/* 	Title           : College Physics
	   	ISBN            : 1002-STUDY-222
	   	Author          : Paul Peter
	   	Price           : USD 3200.00

	   	Title           : Dear John
	   	ISBN            : 1004-NOVEL-444
	   	Author          : Nicholas Sparks
	   	Price           : USD 950.00 */

	book = Book{}
	db.First(&book, "id = ?", "5")
	fmt.Println("----- Get book of id 5 -----")
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)
	/*	Title           : Harry Potter I
		ISBN            : 1005-KIDS-555
		Author          : J. K. Rowling
		Price           : USD 1430.00 */

	// Get all records - SELECT * FROM books;
	db.Find(&books)
	result := db.Find(&books)
	fmt.Println("----- Get all the books -----")
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	// returns found records count, equals `len(books)`
	fmt.Println(result.RowsAffected, "number of records found.")
	fmt.Println(result.Error, "Eroor")

}
