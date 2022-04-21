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
	book2 := Book{ISBN: "1002-STUDY-222", Title: "College Physics", Price: 3200.00, Author: "Paul Peter"}
	book3 := Book{ISBN: "1003-KIDS-333", Title: "The Great Gatsby", Price: 730.00, Author: "F. Scott Fitzgerald"}
	book4 := Book{ISBN: "1004-NOVEL-444", Title: "The Best of Me", Price: 950.00, Author: "Nicholas Sparks"}
	book5 := Book{ISBN: "1005-KIDS-555", Title: "Harry Potter I", Price: 1430.00, Author: "J. K. Rowling"}
	book6 := Book{ISBN: "1006-KIDS-666", Title: "Harry Potter II", Price: 1430.00, Author: "J. K. Rowling"}

	books = []Book{book1, book2, book3, book4, book5, book6}
	db.Create(&books)

	fmt.Println("select * form books order by price desc")
	db.Order("price desc ").Find(&books)
	for _, v := range books {
		fmt.Printf("Title\t: %v \t Price\t: %.2f\n", v.Title, v.Price)
	}

	// Multiple orders
	fmt.Println("\nselect * books order by price desc author")
	db.Order("author desc").Order("price").Find(&books)
	for _, v := range books {
		fmt.Printf("Title\t: %v \t Price\t: %.2f\n", v.Title, v.Price)
	}

	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	/* db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{}) */

}
