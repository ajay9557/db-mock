package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id      int
	name    string
	age     int
	address string
}

func Read(db *sql.DB) (int64, error) {
	rows, err := db.Query("select id from users")
	count := 0

	if err != nil {
		return 0, errors.New("couldn't fetch records")
	}

	for rows.Next() {
		count++
	}

	return int64(count), nil
}

func FetchRecords(db *sql.DB, id int) (User, error) {
	sqlRows := db.QueryRow("select id from users where id = ?", id)

	if sqlRows.Err() != nil {
		return User{}, sql.ErrNoRows
	}

	v := User{}
	sqlRows.Scan(&v.id, &v.name, &v.age, &v.address)

	return v, nil
}

func InsertRecord(db *sql.DB, age int, name, address string) (int64, error) {
	result, err := db.Exec("insert into users(name, age, address) values(?, ?, ?)", name, age, address)

	if err != nil {
		return 0, errors.New("Couldn't insert record")
	}

	i, _ := result.RowsAffected()

	return i, nil
}

func UpdateRecord(db *sql.DB, name string, id int) (int64, error) {

	var result sql.Result
	var err error

	if name != "" {
		result, err = db.Exec("update users set name = ? where id = ?", name, id)
		if err != nil {
			return 0, errors.New("couldn't update record")
		}
	}

	i, _ := result.RowsAffected()

	return i, nil
}

func DeleteRecord(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from users where id = ?", id)
	if err != nil {
		return 0, errors.New("couldn't delete record")
	}

	i, _ := result.RowsAffected()
	return i, nil
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, _ := Read(db)
	fmt.Print("TOtal: ", rows)

	_, err = FetchRecords(db, 1)
	if err != nil {
		fmt.Println(err)
	}
}
