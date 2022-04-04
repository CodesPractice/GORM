package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type Person struct {
	ID        string `gorm:"default:uuid_generate_v3()"` // db func
	FirstName string
	LastName  string
	Age       uint8
	FullName  string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection Error!")
	}
	db.AutoMigrate(&CreditCard{}, &User{}, &Person{})

	/* no := CreditCard{
		Number: "411111111111",
	}
	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: no,
	})

	var u User
	db.Find(&u)

	fmt.Println(u.ID, u.Name, u.CreditCard.Number) */

	per := Person{
		ID:        "1001",
		FirstName: "Chandler",
		LastName:  "Bing",
	}
	db.Create(&per)
	db.Find(&per)
	fmt.Println(per.ID, per.FullName)

}
