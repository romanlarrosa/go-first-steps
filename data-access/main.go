package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	fmt.Println("Albums found: ")
	fmt.Fprintln(w, "ID\tTitle\tArtist\tPrice")
	fmt.Fprintln(w, "--\t-----\t------\t-----")
	for _, album := range albums {
		fmt.Fprintf(w, "%d\t%s\t%s\t%f\n", album.ID, album.Title, album.Artist, album.Price)
	}
	w.Flush()
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}
