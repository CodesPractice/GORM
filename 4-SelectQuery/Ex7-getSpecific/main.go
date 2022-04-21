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

	var user User
	//var users []User

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failiure!")
	}

	/* db.AutoMigrate(&User{})

	u1 := User{Name: "Andrew", Age: 36, Country: "New Zealand"}
	u2 := User{Name: "Rechal", Age: 27, Country: "Australia"}
	u3 := User{Name: "Monika", Age: 18, Country: "Canada"}
	u4 := User{Name: "Soniya", Age: 68, Country: "Bahamas"}
	u5 := User{Name: "Joe", Age: 29, Country: "USA"}
	u6 := User{Name: "Ross", Age: 28, Country: "USA"}

	users = []User{u1, u2, u3, u4, u5, u6}
	db.Create(&users) */

	fmt.Println("Select * From Users where name = 'Joe'	-> FIND")
	db.Find(&user, "name=?", "Joe")
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)

	fmt.Println()
	user = User{}
	fmt.Println("Select * From Users where name = 'Joe'	-> STRUCT")
	db.Find(&user, User{Name: "Joe"})
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)

	fmt.Println()
	user = User{}
	fmt.Println("Select * From Users where name = 'Joe'	-> MAP")
	db.Find(&user, map[string]interface{}{"Name": "Joe"})
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println("-----------------------------------")

	fmt.Println()
	user = User{}
	fmt.Println("Select * From Users where name = 'Joe'	-> WHERE")
	db.Where("name =?", "joe").Find(&user)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println("-----------------------------------")
}
