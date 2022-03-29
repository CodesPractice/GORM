package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}

type Book struct {
	gorm.Model
	Title     string
	ISBN      string
	NoOfPages uint
	Publisher string `gorm:"default:HarperCollins"`
	Price     float64
	Author    Author `gorm:"embedded"`
	Country   string `gorm:"default:USA"`
	Language  string `gorm:"default:English"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection Error!")
	}

	db.AutoMigrate(&Author{}, &Book{})

	// conunty and language are not inserted
	// default values are passed to the database

	/* 	author1 := Author{
	   		FirstName: "Charles",
	   		LastName:  "Dickens",
	   	}
	   	book1 := Book{
	   		Title:     "A Tale of Two Cities",
	   		ISBN:      "1503219704",
	   		NoOfPages: 290,
	   		Publisher: "Public Domain Books",
	   		Price:     350,
	   		Author:    author1,
	   	}
	   	db.Create(&book1) */

	var b Book
	db.First(&b)
	fmt.Printf(" Title \t\t: %v\n ISBN \t\t: %v\n Publisher \t: %v\n Author \t: %v\n Country \t: %v\n Language \t: %v\n pages \t\t: %v\n Price \t\t: USD %0.2f\n", b.Title, b.ISBN, b.Publisher, b.Author.FirstName+" "+b.Author.LastName, b.Country, b.Language, b.NoOfPages, b.Price)

}
