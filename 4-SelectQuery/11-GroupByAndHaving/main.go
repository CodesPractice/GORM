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

	/* 	book1 := Book{ISBN: "1001-NOVEL-111", Title: "A Walk to Remember", Price: 800.00, Author: "Nicholas Sparks"}
	   	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter   "}
	   	book3 := Book{ISBN: "1003-KIDS-333", Title: "The Great Gatsby", Price: 730.00, Author: "F. Scott Fitzgerald"}
	   	book4 := Book{ISBN: "1004-NOVEL-444", Title: "The Best of Me", Price: 950.00, Author: "Nicholas Sparks"}
	   	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	   	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	   	books = []Book{book1, book2, book3, book4, book5, book6}
	   	db.Create(&books) */

	fmt.Println("SELECT ISBN, Author, sum(price) as Price FROM `book`  GROUP BY `Author`")
	db.Model(&Book{}).Select("author, ISBN, sum(price) as price").Group("Author").Find(&books)
	for _, v := range books {
		fmt.Printf("ISBN\t: %v \t Author\t: %v \t Total\t: %.2f\n", v.ISBN, v.Author, v.Price)
	}

	fmt.Println()
	fmt.Println("SELECT  Author, sum(price) as Price FROM `book`  where ISBN like '%kids%' GROUP BY `Author`")
	db.Model(&Book{}).Select("author, sum(price) as price").Where("ISBN LIKE ?", "%KIDS%").Group("author").Find(&books)
	for _, v := range books {
		fmt.Printf("Author\t: %v \t Total\t: %.2f\n", v.Author, v.Price)
	}

	fmt.Println()
	fmt.Println("SELECT  Author, sum(price) as Price FROM `book`  GROUP BY `Author` having ISBN like '%kids%'")
	db.Model(&Book{}).Select("author, sum(price) as price").Having("ISBN LIKE ?", "%KIDS%").Group("author").Find(&books)
	for _, v := range books {
		fmt.Printf("Author\t: %v \t Total\t: %.2f\n", v.Author, v.Price)
	}

}
