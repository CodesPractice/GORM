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

	fmt.Println("select * from books where title ='Harry Potter I' OR price > 1300.00")
	db.Where("Title = ?", "Harry Potter I").Or("Price >?", 1300.00).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("select * from books where ISBN: '1005-KIDS-555' OR	Author ='Nicholas Sparks' AND price = 950.00  ->STRUCT")
	db.Where(Book{ISBN: "1005-KIDS-555"}).Or(Book{Author: "Nicholas Sparks", Price: 950}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("select * from books where author = 'Paul Peter' OR price = 430.00  ->MAP")
	db.Where(map[string]interface{}{"Author": "Paul Peter"}).Or(map[string]interface{}{"Price": 430}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

}
