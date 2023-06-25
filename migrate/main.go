package main

import (
	"log"

	"github.com/TulioMeran/gin_example/initializer"
	"github.com/TulioMeran/gin_example/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.DbConnection()
}

func main() {
	err := initializer.DB.AutoMigrate(&models.Ticket{})

	if err != nil {
		log.Fatalln("Error migrating database: " + err.Error())
		return
	}

	log.Println("Migrating success!!!")
}
