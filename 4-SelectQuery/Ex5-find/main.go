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

	// Get by primary key if it were a non-integer type
	/* fmt.Println("SELECT * FROM users WHERE id = 'string_primary_key'")
	db.First(&user, "id = ?", " ")
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println() */

	// Get by primary key if it were a non-integer type
	fmt.Println("SELECT * FROM users WHERE Name = 'Rechal'")
	db.First(&user, "Name = ?", "Rechal")
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println()

	user = User{}
	db.Find(&users, "name <> ? AND age > ?", "Rechal", 27)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	user = User{}
	db.Find(&users, User{Age: 18})
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	// SELECT * FROM users WHERE age = 20;

	// Map
	user = User{}
	db.Find(&users, map[string]interface{}{"age": 20})
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	// SELECT * FROM users WHERE age = 20;

}
