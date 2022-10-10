package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(url string) {
	var err error
	fmt.Println("Connect to DB")
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("DB connect error")
		log.Fatalf("%v", err)
	}

	fmt.Println("DB connect succeed")
}
