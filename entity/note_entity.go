package entity

import (
	"time"
)

type Note struct {
	Id          uint      `gorm:"primaryKey:autoIncrement"`
	Title       string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Description string
	UserId      int64 `gorm:"not null"`
}
