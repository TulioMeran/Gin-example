package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DbConnection() {
	HOST_NAME := os.Getenv("HOST_DB")
	USER_NAME := os.Getenv("USER_DB")
	PASSWORD := os.Getenv("PASSWORD_DB")
	DB_NAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("PORT_DB")

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", HOST_NAME, USER_NAME, PASSWORD, DB_NAME, PORT)

	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connection with the database: " + err.Error())
		return
	}

	log.Println("DB connected!!")

}
