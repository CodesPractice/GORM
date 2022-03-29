package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	Name    string
	Email   string
	Address string
}

var personInfo []Person

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failure!")
	}
	db.AutoMigrate(&Person{})

	per := getPersonInfo()
	db.Create(&per)

	db.Find(&personInfo) // find Person data form database and store in personInfor slice variable
	for _, v := range personInfo {
		fmt.Println(v.Name, v.Address, v.Email)
	}
}

func getPersonInfo() []Person {

	var name, email, add string
	var exit string

	for {
		fmt.Println("Input following infomation of the person.")
		fmt.Print("Name :")
		fmt.Scan(&name)
		fmt.Print("Email : ")
		fmt.Scan(&email)
		fmt.Print("Address : ")
		fmt.Scan(&add)

		per := Person{Name: name, Email: email, Address: add}
		personInfo = append(personInfo, per)

		fmt.Println("Do you want add another person infor (y/n)")
		fmt.Scan(&exit)
		if exit != "y" {
			goto out
		}
	}
out:
	return personInfo
}
