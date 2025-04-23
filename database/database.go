package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gmalheiro/golang-playground-txdb/internal/configs"
)

var conn *sql.DB

func Connect() error {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.Get().DbHost,
		configs.Get().DbPort,
		configs.Get().DbUser,
		configs.Get().DbPassword,
		configs.Get().DbName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		log.Printf("err while pinging to db err: %s", err.Error())
		return err
	}

	conn = db

	return nil
}

func GetConnection() *sql.DB {
	return conn
}
