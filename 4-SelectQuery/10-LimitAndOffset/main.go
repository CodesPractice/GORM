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

	book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter   "}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "The Great Gatsby", Price: 730.00, Author: "F. Scott Fitzgerald"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "The Best of Me", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	books = []Book{book1, book2, book3, book4, book5, book6}
	db.Create(&books)

	fmt.Println("select * form books limit --> retrive only first 2 records")
	db.Limit(2).Find(&books)
	for _, v := range books {
		fmt.Printf("ISBN\t: %v \t Title\t: %v \t Author\t: %v\n", v.ISBN, v.Title, v.Author)
	}

	fmt.Println("\nselect * form books limit 3		--> retrive only first 3 records")
	fmt.Println("select * form books // without limit --> retrive all the records")
	var books1 []Book
	var books2 []Book
	db.Limit(3).Find(&book1).Limit(-1).Find(&book2)
	for _, v := range books1 {
		fmt.Printf("ISBN\t: %v \t Title\t: %v \t Author\t: %v\n", v.ISBN, v.Title, v.Author)
	}
	fmt.Println()
	for _, v := range books2 {
		fmt.Printf("ISBN\t: %v \t Title\t: %v \t Author\t: %v\n", v.ISBN, v.Title, v.Author)
	}

	// skip 4 records from the begin of the list
	fmt.Println("\nselect * form books Offset 4	--> skip first 4 records")
	db.Offset(4).Find(&books)
	for _, v := range books {
		fmt.Printf("ISBN\t: %v \t Title\t: %v \t Author\t: %v\n", v.ISBN, v.Title, v.Author)
	}

	fmt.Println("\nselect * form books Offset 2 and limit 3	--> skip first 3 records and retrive only next 3 records")
	db.Limit(3).Offset(2).Find(&books)
	for _, v := range books {
		fmt.Printf("ISBN\t: %v \t Title\t: %v \t Author\t: %v\n", v.ISBN, v.Title, v.Author)
	}

}
