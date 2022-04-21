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
	var users []User

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failiure!")
	}

	db.AutoMigrate(&User{})

	u1 := User{Name: "Andrew", Age: 36, Country: "New Zealand"}
	u2 := User{Name: "Rechal", Age: 27, Country: "Australia"}
	u3 := User{Name: "Monika", Age: 18, Country: "Canada"}
	u4 := User{Name: "Soniya", Age: 68, Country: "Bahamas"}
	u5 := User{Name: "Joe", Age: 29, Country: "USA"}
	u6 := User{Name: "Ross", Age: 28, Country: "USA"}

	users = []User{u1, u2, u3, u4, u5, u6}
	db.Create(&users)

	fmt.Println("Select * From users where name = 'Joe'	")
	db.Find(&user, "name=?", "Joe")
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println()

	fmt.Println("Select * From users where country  <> 'USA'")
	db.Find(&users, "country <>?", "USA")
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	user = User{}
	fmt.Println("Select * From users where name = 'Joe'  -> WHERE")
	db.Where("name = ?", "Joe").Find(&user)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)
	fmt.Println()

	users = []User{}
	fmt.Println("Select * From users where country <> 'USA'  -> WHERE")
	db.Where("country <> ?", "USA").Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	users = []User{}
	fmt.Println("Select * From users where age between 25 and 35  -> WHERE")
	db.Where("age BETWEEN ? AND ?", 25, 35).Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	users = []User{}
	fmt.Println("Select * From users where age > 55  -> WHERE")
	db.Where("age > ?", 55).Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	users = []User{}
	fmt.Println("Select * From users where country in ('USA','Australia')  -> WHERE")
	db.Where("country IN ?", []string{"USA", "Australia"}).Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

	users = []User{}
	fmt.Println("Select * From users where name LIKE %a -> WHERE")
	db.Where("name LIKE ?", "%a").Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}
	fmt.Println()

}
