package controllers

import (
	"log"
	"net/http"

	"github.com/TulioMeran/gin_example/dto"
	"github.com/TulioMeran/gin_example/initializer"
	"github.com/TulioMeran/gin_example/models"
	"github.com/gin-gonic/gin"
)

func GetAllTickets(c *gin.Context) {

	var tickets []models.Ticket

	result := initializer.DB.Find(&tickets)

	if result.Error != nil {
		log.Fatalln("Error happen getting all tickets: " + result.Error.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error getting all tickets",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func PostTicket(c *gin.Context) {

	var b dto.TicketDto

	if err := c.Bind(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(b.TicketNumber) < 1 {
		log.Println(b.TicketNumber)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ticketNumber is required",
		})
		return
	}

	if len(b.Notes) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "notes is required",
		})
		return
	}

	var newTicket models.Ticket

	newTicket.TicketNumber = b.TicketNumber
	newTicket.Notes = b.Notes

	result := initializer.DB.Create(&newTicket)

	if result.Error != nil {
		log.Fatalln("Error happened creating ticket: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error happened creating ticket.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newTicket,
	})

}

func UpdateTicket(c *gin.Context) {
	id := c.Query("id")

	var b dto.TicketDto

	c.Bind(&b)

	if len(id) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param is required",
		})
		return
	}

	var t models.Ticket
	result := initializer.DB.First(&t, id)

	if result.Error != nil {
		log.Println("Error happened getting single ticket: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error happened getting single ticket",
		})
		return
	}

	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ticket not exists",
		})
		return
	}

	if len(b.TicketNumber) > 1 {
		t.TicketNumber = b.TicketNumber
	}

	if len(b.Notes) > 1 {
		t.Notes = b.Notes
	}

	result = initializer.DB.Save(&t)

	if result.Error != nil {
		log.Println("Error happened updating ticket: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error happened updating ticket",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": t,
	})

}

func DeleteTicket(c *gin.Context) {
	id := c.Query("id")

	if len(id) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param is required",
		})
		return
	}

	var t models.Ticket
	result := initializer.DB.First(&t, id)

	if result.Error != nil {
		log.Println("Error happened getting single ticket: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error happened getting single ticket",
		})
		return
	}

	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ticket not exists",
		})
		return
	}

	result = initializer.DB.Unscoped().Delete(&t)

	if result.Error != nil {
		log.Println("Error happened deleting ticket: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error happened deleting ticket",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "ticket deleted",
	})

}
