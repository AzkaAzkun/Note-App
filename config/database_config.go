package config

import (
	"fmt"
	"os"

	"github.com/jeypc/go-jwt-mux/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DBUSER := os.Getenv("DB_user")
	DBPASS := os.Getenv("DB_pass")
	DBHOST := os.Getenv("DB_host")
	DBPORT := os.Getenv("DB_port")
	DBNAME := os.Getenv("DB_name")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBUSER, DBPASS, DBHOST, DBPORT, DBNAME)

	gormDB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database " + err.Error())
	}
	migrate(gormDB)
	DB = gormDB
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Note{})
}
