package driver

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func ConnectDB() *sql.DB {

	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(fmt.Sprintf("Database se connection nahi bana: %s", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Database is not able to ping : %s", err.Error()))
	}

	fmt.Println("Database connected successfully")

	DB = db
	return db
}
