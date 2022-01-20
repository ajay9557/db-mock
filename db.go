package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id      int64
	name    string
	age     int
	address string
}

var (
	id      int
	name    string
	age     string
	address string
)

var users []User

func Read(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")

	if err != nil {
		return []User{}, errors.New("couldn't fetch records")
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &address)

		if err != nil {
			return []User{}, err
		}

		convAge, _ := strconv.Atoi(age)
		u := User{id: int64(id), name: name, age: convAge, address: address}
		users = append(users, u)
	}

	return users, nil
}

func ReadById(db *sql.DB, id int) (User, error) {
	sqlRows := db.QueryRow("select * from users where id = ?", id)

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

	id, _ := result.LastInsertId()

	return id, nil
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

	i, _ := result.LastInsertId()
	return i, nil
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	InsertRecord(db, 21, "Ridhdhish", "India")
	InsertRecord(db, 18, "Naruto", "Japan")

	UpdateRecord(db, "Rid", 1)
	data, err := ReadById(db, 1)
	fmt.Println(data)

	DeleteRecord(db, 1)

	rows, _ := Read(db)
	fmt.Print("All Records: ", rows)

	if err != nil {
		fmt.Println(err)
	}
}
