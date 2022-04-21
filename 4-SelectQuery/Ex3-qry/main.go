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

	// // Struct
	fmt.Println("------Struct------	\nSELECT * FROM users WHERE name ='Andrew' and age = 36")
	db.Where(&User{Name: "Andrew", Age: 36}).First(&user)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)

	user = User{}
	fmt.Println("\n------Map------ \nSELECT * FROM users WHERE name ='Rechal' and age = 27")
	db.Where(map[string]interface{}{"Name": "Rechal", "Age": 27}).Find(&user)
	fmt.Printf("%d\t%s\t%d\t%s\n", user.ID, user.Name, user.Age, user.Country)

	users = []User{}
	fmt.Println("\n------slice of ID s------ \nSELECT * FROM users WHERE Id in (1,3)")
	db.Where([]int{1, 3}).Find(&users)
	for _, v := range users {
		fmt.Printf("%d\t%s\t%d\t%s\n", v.ID, v.Name, v.Age, v.Country)
	}

}
