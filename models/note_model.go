package models

import (
	"time"
)

type Note struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string
}
