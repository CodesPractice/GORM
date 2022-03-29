package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Owner struct {
	OwnerName string
	Age       uint
	Country   string `gorm:"default:USA"`
}

type Pet struct {
	Name  string
	Dob   time.Time
	Breed string
	Owner Owner `gorm:"embedded"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Connection Failure!")
	}
	db.AutoMigrate(&Owner{}, &Pet{})

	owner1 := Owner{OwnerName: "Mark Antony", Age: 55, Country: "Australia"}
	db.Create(&owner1)
	pet1 := Pet{Name: "Mithu", Dob: time.Now(), Breed: "Pomaranian", Owner: owner1}
	db.Create(&pet1)

	var ow Owner
	var pe Pet

	db.First(&ow)
	db.First(&pe)

	fmt.Println(ow.Country)
	fmt.Println(pe.Name)

	//----------------------------

	owner2 := Owner{OwnerName: "Anna Marry", Age: 33, Country: "Emirates"}
	db.Select("Name", "Age").Create(&owner2)

	owner3 := Owner{OwnerName: "Pheebe Buffer", Age: 38, Country: "Canada"}
	db.Omit("Age").Create(&owner3)

	pet2 := Pet{Name: "Kuky", Dob: time.Now(), Breed: "Natural", Owner: owner2}

	pet3 := Pet{Name: "Teena", Dob: time.Now(), Breed: "Dalmation", Owner: owner3}

	db.Create(&pet2)
	db.Create(&pet3)

	var dogs []Pet

	db.Find(&dogs)
	fmt.Println("Pet\t Breed\t\t Owner\t\t Country\t Age")
	for _, v := range dogs {
		fmt.Println(v.Name, "\t", v.Breed, "\t", v.Owner.OwnerName, "\t", v.Owner.Country, "\t", v.Owner.Age)
	}

}
