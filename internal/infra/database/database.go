package database

import (
	"beyond/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Falha ao conectar ao Postgres: ", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Executa a migração
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Falha ao realizar a migração: ", err)
	}

	log.Println("Migração concluída: Tabela 'users' pronta!")

	DB = database
}
