package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func New(d *sql.DB) {
	db = d
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

/*type Bowler struct{
	Name string
	Age int64
	Wickets int64
}*/

func main() {
	cfg := mysql.Config{
		User:   "nayani",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	err = updateAlbum(2)
	if err != nil {
		fmt.Printf("Update failed")
	}

	albums, err := albumsByArtist("John Coltrne")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums fund: %v\n", albums)

	err = delete(3)
	if err != nil {
		fmt.Printf("Update failed")
	}

	alb, err := albumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
}

func updateAlbum(id int) error {
	_, err := db.Exec("Update Album set Price=79.99 where id=?", id)
	if err != nil {
		//fmt.Println("Update failed")
		return err
	}
	return nil
}

func delete(id int) error {
	_, err := db.Exec("Delete From Album where id=?", id)
	if err != nil {
		//fmt.Println("Update failed")
		return err
	}
	return nil
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsbyArtist %q: %v", name, err)

	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsbyArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsbyArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id=?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("no albums found")
		}
		return alb, fmt.Errorf("albumByID:%d: %v", id, err)

	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	//fmt.Println(alb)
	result, err := db.Exec("INSERT INTO album(title,artist,price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addA;bum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ehey")
	}
	return id, nil
}
