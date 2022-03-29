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

	fmt.Println("---------SELECT * FROM books WHERE Author = 'Nicholas Sparks' ORDER BY id LIMIT 1;")
	db.Where("Author = ?", "Nicholas Sparks").First(&book)
	fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", book.Title, book.ISBN, book.Author, book.Price)

	/* 	Title           : A Walk to Remember
	   	ISBN            : 1001-NOVEL-111
	   	Author          : Nicholas Sparks
	   	Price           : USD 800.00     */

	fmt.Println("---------SELECT * FROM books WHERE Author equl to 'Nicholas Sparks' ORDER BY id")
	db.Where("Author = ?", "Nicholas Sparks").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/*	Title           : A Walk to Remember
		ISBN            : 1001-NOVEL-111
		Author          : Nicholas Sparks
		Price           : USD 800.00

		Title           : Dear John
		ISBN            : 1004-NOVEL-444
		Author          : Nicholas Sparks
		Price           : USD 950.00 */

	fmt.Println("---------SELECT * FROM books WHERE Author not equal to 'Nicholas Sparks' ORDER BY id")
	db.Where("Author <> ?", "Nicholas Sparks").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/*	Title           : College Physics
		ISBN            : 1002-STUDY-222
		Author          : Paul Peter
		Price           : USD 3200.00

		Title           : PINOKIO
		ISBN            : 1003-KIDS-333
		Author          : Carlo Collodi
		Price           : USD 430.00

		Title           : Harry Potter I
		ISBN            : 1005-KIDS-555
		Author          : J. K. Rowling
		Price           : USD 1430.00

		Title           : Harry Potter II
		ISBN            : 1006-KIDS-666
		Author          : J. K. Rowling
		Price           : USD 1430.00 */

	fmt.Println("---------IN - SELECT * FROM users WHERE name IN (Carlo Collodi, Paul Peter);")
	db.Where("Author IN ?", []string{"Carlo Collodi", "Paul Peter"}).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/*	Title           : College Physics
		ISBN            : 1002-STUDY-222
		Author          : Paul Peter
		Price           : USD 3200.00

		Title           : PINOKIO
		ISBN            : 1003-KIDS-333
		Author          : Carlo Collodi
		Price           : USD 430.00
	*/

	fmt.Println("---------LIKE - SELECT * FROM books WHERE Author LIKE '%ll%';")
	db.Where("Author LIKE ?", "%ll%").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/* 	Title           : PINOKIO
	   	ISBN            : 1003-KIDS-333
	   	Author          : Carlo Collodi
	   	Price           : USD 430.00 */

	fmt.Println("---------AND - SELECT * FROM books WHERE Author = 'J. K. Rowling' AND ISBN = 1005-KIDS-555;")
	db.Where("Author = ? AND ISBN = ?", "J. K. Rowling", "1005-KIDS-555").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/*	Title           : Harry Potter I
		ISBN            : 1005-KIDS-555
		Author          : J. K. Rowling
		Price           : USD 1430.00 */

	fmt.Println("---------SELECT * FROM books WHERE updated_at > 2022-03-29 00:00:00")
	db.Where("updated_at > ?", "2022-03-29 00:00:00").Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}

	fmt.Println("---------SELECT * FROM books WHERE Price BETWEEN 800 AND 1000 ")
	db.Where("Price BETWEEN ? AND ?", 800, 1000).Find(&books)
	for _, v := range books {
		fmt.Printf("Title \t\t: %v\nISBN \t\t: %v\nAuthor \t\t: %v\nPrice \t\t: USD %0.2f\n\n", v.Title, v.ISBN, v.Author, v.Price)
	}
	/*
		Title           : A Walk to Remember
		ISBN            : 1001-NOVEL-111
		Author          : Nicholas Sparks
		Price           : USD 800.00

		Title           : PINOKIO
		ISBN            : 1003-KIDS-333
		Author          : Carlo Collodi
		Price           : USD 430.00

		Title           : Dear John
		ISBN            : 1004-NOVEL-444
		Author          : Nicholas Sparks
		Price           : USD 950.00
	*/
}
