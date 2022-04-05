package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=172.23.0.2 user=root password=root dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}
}
