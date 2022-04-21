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

func main() {

	//var user User

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failiure!")
	}

	db.AutoMigrate(&User{})

	u1 := User{Name: "Andrew", Age: 36, Country: "New Zealand"}
	u2 := User{Name: "Rechal", Age: 27, Country: "Australia"}
	u3 := User{Name: "Monika", Age: 18, Country: "Canada"}
	u4 := User{Name: "Soniya", Age: 68, Country: "Bahamas"}

	users := []User{u1, u2, u3, u4}
	db.Create(&users)

	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	db.Where(&User{Name: "Andrew"}, "name", "Age").Find(&users)

	db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}

}
