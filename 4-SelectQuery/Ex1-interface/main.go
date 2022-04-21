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

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failiure!")
	}

	db.AutoMigrate(&User{})

	u1 := User{Name: "Andrew", Age: 36, Country: "New Zealand"}
	u2 := User{Name: "Rechal", Age: 27, Country: "Australia"}
	u3 := User{Name: "Monika", Age: 22, Country: "Canada"}
	u4 := User{Name: "Soniya", Age: 68, Country: "Bahamas"}

	users := []User{u1, u2, u3, u4}
	db.Create(&users)

	// works because model is specified using `db.Model()`
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	result := map[string]interface{}{}
	db.Model(&User{}).Last(&result)
	fmt.Println(result)
	//map[age:68 country:Bahamas created_at:2022-04-20 11:35:43.8902604 +0530 +0530 deleted_at:<nil> id:8 name:Soniya updated_at:2022-04-20 11:35:43.8902604 +0530 +0530]

	db.Table("users").Take(&result)
	fmt.Println(result)
	//map[age:36 country:New Zealand created_at:2022-04-20 11:33:55.682958 +0530 +0530 deleted_at:<nil> id:1 name:Andrew updated_at:2022-04-20 11:33:55.682958 +0530 +0530]

}
