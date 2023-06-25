package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	TicketNumber string    `gorm:"not null" json:"ticketNumber"`
	Notes        string    `json:"notes"`
	RegisterTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"registerTime"`
}
