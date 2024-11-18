package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {

	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	var err error
	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	fmt.Println("Database connected successfully!")
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("failed to close database connection", err)
	}
}
