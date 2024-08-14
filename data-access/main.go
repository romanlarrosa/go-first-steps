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
	godotenv.Load()

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

	// Get albums by artist
	artist := "John Coltrane"
	fmt.Printf("Querying albums by artist: %s...\n", artist)
	albums, err := albumsByArtist(artist)
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

	// Get album by id
	albumId := 3
	fmt.Printf("Querying album by ID: %d...\n", albumId)
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Album found:")
	fmt.Fprintln(w, "ID\tTitle\tArtist\tPrice")
	fmt.Fprintln(w, "--\t-----\t------\t-----")
	fmt.Fprintf(w, "%d\t%s\t%s\t%f\n", alb.ID, alb.Title, alb.Artist, alb.Price)
	w.Flush()

	//Add an album
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
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

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
