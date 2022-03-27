package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       uint
	Address   string
}

type Author struct {
	ID    int
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

/* type User struct {
	a string `gorm:"<-:create"`          // allow read and create
	b string `gorm:"<-:update"`          // allow read and update
	c string `gorm:"<-"`                 // allow read and write (create and update)
	d string `gorm:"<-:false"`           // allow read, disable write permission
	e string `gorm:"->"`                 // readonly (disable write permission unless it configured)
	f string `gorm:"->;<-:create"`       // allow read and create
	g string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	h string `gorm:"-"`                  // ignore this field when write and read with struct
	i string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	j string `gorm:"-:migration"`        // ignore this field when migrate with struct
} */

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Databese connection failure!")
	}
	fmt.Println("Successfully conncted")

	db.AutoMigrate(&Person{}, &Author{}, &Blog{})

	blg1 := Blog{
		Author:  Author{Name: "Todd MacLead", Email: "temc@yahoo.com"},
		Upvotes: 22,
	}
	db.Create(&blg1)
}
