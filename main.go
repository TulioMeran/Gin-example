package main

import (
	"github.com/TulioMeran/gin_example/controllers"
	"github.com/TulioMeran/gin_example/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.DbConnection()

}

func main() {
	r := gin.Default()

	ticketController := r.Group("/api/ticket")
	{
		ticketController.GET("/", controllers.GetAllTickets)
		ticketController.POST("/", controllers.PostTicket)
		ticketController.PUT("/", controllers.UpdateTicket)
		ticketController.DELETE("/", controllers.DeleteTicket)
	}

	r.Run()
}
