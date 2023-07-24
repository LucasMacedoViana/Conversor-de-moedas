package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

// OpenConnection abre a conexão com o banco de dados
func OpenConnection() *gorm.DB {
	host := os.Getenv("host")
	dbname := os.Getenv("dbname")
	port := os.Getenv("port_banco")
	user := os.Getenv("user")
	password := os.Getenv("password")
	str := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s", host, dbname, port, user, password)

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}
	return database
}

// CloseConnection fecha a conexão com o banco de dados
func CloseConnection(connection *gorm.DB) {
	db, err := connection.DB()
	if err != nil {
		log.Fatalln(err)
	}
	db.Close()
}
