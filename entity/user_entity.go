package entity

type User struct {
	Id       uint   `gorm:"primaryKey:autoIncrement"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Notes    []Note `gorm:"foreignKey:UserId"`
}
