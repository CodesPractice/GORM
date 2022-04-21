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

	//var book Book
	var books []Book

	/* book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter"}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "PINOKIO", Price: 430.00, Author: "Carlo Collodi"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "Dear John", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	books = []Book{book1, book2, book3, book4, book5, book6}
	db.Create(&books) */

	fmt.Println("Select * from books where author = 'J. K. Rowling'")
	db.Find(&books, "author =?", "J. K. Rowling")
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author <> 'J. K. Rowling'")
	db.Not("author =?", "J. K. Rowling").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author <> 'Carlo Collodi' -> MAPS")
	db.Not(map[string]interface{}{"Author": "Carlo Collodi"}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author <> 'J. K. Rowling' -> STRUCT")
	db.Not(Book{Author: "J. K. Rowling"}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author not in ('J. K. Rowling','Nicholas Sparks') -> MAP")
	db.Not(map[string]interface{}{"author": []string{"J. K. Rowling", "Nicholas Sparks"}}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author not in ('J. K. Rowling','Nicholas Sparks' limit 1) -> MAP")
	db.Not(map[string]interface{}{"Author": []string{"J. K. Rowling", "Nicholas Sparks"}}).First(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("Select * from books where author <> 'Nicholas Sparks' and price <> 1430.00 and ISBN <> '1002-STUDY-222' -> STRUCT")
	db.Not(Book{Author: "Nicholas Sparks", Price: 1430.00, ISBN: "1002-STUDY-222"}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

}
