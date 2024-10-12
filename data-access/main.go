package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Album struct {
	id     int64
	title  string
	artist string
	price  float32
}

func main() {
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumByID(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", album)
}

func getDBHandle() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	return db, err
}

func albumByID(id int64) (Album, error) {
	var alb Album
	db, err := getDBHandle()

	if err != nil {
		log.Fatal("Error connecting")
	}

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&alb.id, &alb.title, &alb.artist, &alb.price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}

		return alb, fmt.Errorf("albumsById: %d %v", id, err)
	}

	return alb, nil
}

func albumsByArtist(name string) ([]Album, error) {
	db, err := getDBHandle()

	if err != nil {
		log.Fatal("Error connecting")
	}

	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist: %q, %v", name, err)
	}
	defer rows.Close()

	// Iterate rows
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.id, &alb.title, &alb.artist, &alb.price); err != nil {
			return nil, fmt.Errorf("albumsByArtist: %q, %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist: %q %v", name, err)
	}

	if err = db.Close(); err != nil {
		return nil, err
	}

	return albums, nil
}
