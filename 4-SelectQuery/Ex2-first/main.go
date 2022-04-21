package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Age     uint
	Country string
}

type Book struct {
	Title  string
	Pages  uint
	Author string
}

func main() {

	var user User
	//var book Book

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failiure!")
	}

	db.AutoMigrate(&User{}, &Book{})

	u1 := User{Name: "Andrew", Age: 36, Country: "New Zealand"}
	u2 := User{Name: "Rechal", Age: 27, Country: "Australia"}
	u3 := User{Name: "Monika", Age: 22, Country: "Canada"}
	u4 := User{Name: "Soniya", Age: 68, Country: "Bahamas"}

	users := []User{u1, u2, u3, u4}
	db.Create(&users)

	b1 := Book{"Philosopher's Stone", 1200, "J. K. Rowling"}
	b2 := Book{"Chamber of Secrets", 980, "J. K. Rowling"}
	b3 := Book{"The Adventures of Tintin", 1350, "Georges Prosper Remi"}
	b4 := Book{"The Adventures of Asterix", 800, "Albert Uderzo"}

	books := []Book{b1, b2, b3, b4}
	db.Create(&books)

	fmt.Println("SELECT * FROM users WHERE id = 3")
	db.First(&user, 3)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println()
	//	3       Monika  22      Canada

	fmt.Println("SELECT * FROM users WHERE id IN (1,2,3)")
	db.Find(&users, []int{1, 3, 2})
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	/*
		1       Andrew  36      New Zealand
		2       Rechal  27      Australia
		3       Monika  22      Canada
	*/

	// get the first matched record
	fmt.Println("SELECT * FROM users WHERE name = 'Rechal' ORDER BY id LIMIT 1;")
	user = User{} //crean the variable
	db.Where("Name = ?", "Rechal").First(&user)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println()
	//	2       Rechal  27      Australia
}
