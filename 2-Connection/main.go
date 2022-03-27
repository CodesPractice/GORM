// step 1 - run go mod init Connection

// step 2 - go get -u gorm.io/gorm
//			go get -u gorm.io/driver/sqlite

// once completed the code
//			go run main.go

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database Connection Failiur!")
	}
	fmt.Println("Successfully Connected to the database")
}
