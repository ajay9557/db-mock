package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlDB struct {
	db *sql.DB
}

type User struct {
	id        int
	firstName string
	lastName  string
	age       int
}

func connectDb() (*sql.DB, error) {
	connection, error := sql.Open("mysql", "vicky:root@/test")
	if error != nil {
		return nil, errors.New("ERROR IN LOG-IN DB")
	}
	return connection, nil
}

func (db *SqlDB) AddUser(id int, firstName string, lastName string, age int) error {
	_, error := db.db.Exec("insert into userdemo1(id, firstname, lastname, age) values(?, ?, ?, ?)", id, firstName, lastName, age)

	if id > 100 {
		return errors.New("INVALID ID")
	}

	if error != nil {
		return errors.New("FAILED TO ADD USER")
	}
	return nil
}

func (db *SqlDB) UpdateUser(id int, firstName string, lastName string, age int) error {
	_, error := db.db.Exec("update userdemo1 set firstname=?, lastname=?, age=? where id=?", firstName, lastName, age, id)
	if error != nil {
		return errors.New("FAILED TO UPDATE USER")
	}
	return nil
}

func (db *SqlDB) GetUser(id int) (User, error) {
	usr := User{}
	rows := db.db.QueryRow("select * from userdemo1 where id=?", id)

	error := rows.Scan(&usr.id, &usr.firstName, &usr.lastName, &usr.age)

	if error != nil {
		return usr, errors.New("PROBLEM IN GETTING A ROW")
	}

	return usr, nil
}

func (db *SqlDB) GetUsers() ([]User, error) {
	userList := make([]User, 10)
	rows, error := db.db.Query("select * from userdemo1")

	if error != nil {
		return userList, errors.New("ERROR IN FETCHING ROW")
	}

	defer rows.Close()

	for rows.Next() {
		usr := User{}
		error := rows.Scan(&usr.id, &usr.firstName, &usr.lastName, &usr.age)
		if error != nil {
			return userList, errors.New("ERROR IN FETCHING ROW")
		}
		userList = append(userList, usr)
	}
	return userList, nil
}

func (db *SqlDB) DeleteUser(id int) error {
	_, error := db.db.Exec("delete from userdemo1 where id=?", id)

	if id > 100 {
		return errors.New("FAILED TO DELETE THE USER")
	}

	if error != nil {
		return errors.New("ID NOT FOUND")
	}
	return nil
}

func main() {
	fmt.Println("Database mocking...")
	connection, _ := connectDb()
	dsc := SqlDB{connection}
	//user, _ := dsc.GetUser(10)
	fmt.Println(dsc.AddUser(11, "Ram", "Sharma", 22))

	//fmt.Println(dsc.DeleteUser(88))
}
