package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Load env vars
	godotenv.Load("./data/.env")

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("MYSQLUSER"),
		Passwd: os.Getenv("MYSQLPASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_PUBLIC_URL"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
