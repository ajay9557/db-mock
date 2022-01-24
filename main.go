package main

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type UserDetails struct {
	Name    string
	Age     int
	Address string
	Delete  bool
}

func UpdateUser(db *sql.DB, user *UserDetails) (err error) {
	updateQ := "update user set " //Update query
	if user.Age > 0 {
		updateQ += "Age=?,"
	}
	if user.Address != "" {
		updateQ += "Address=?,"
	}
	updateQ += "Delete=? "
	updateQ = strings.TrimRight(updateQ, ",")
	if user.Name != "" {
		updateQ += " where Name=?;"
	}

	_, err = db.Exec(updateQ, user.Age, user.Address, user.Delete, user.Name)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func ReadUser(db *sql.DB, name string) (user *UserDetails, err error) {
	ReadQ := "SELECT * FROM user where Name=?;" //Fetch values from table
	user = &UserDetails{}
	rows, err := db.Query(ReadQ, name)
	if err != nil {
		return nil, errors.New("couldn't exec query")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Age, &user.Address, &user.Delete)
		if err != nil {
			return nil, errors.New("couldn't exec query")
		}
	}
	return user, nil
}

func InsertUser(db *sql.DB, user *UserDetails) (err error) {
	// Create and insert values into table
	insertQ := "INSERT INTO user(Name, Age, Address,Del) VALUES(?,?,?,?);" //Insert query
	_, err = db.Exec(insertQ, user.Name, user.Age, user.Address, user.Delete)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func DeleteUser(db *sql.DB) (err error) {
	deleteQ := "DELETE From user where Del=?;"
	_, err = db.Exec(deleteQ, true)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func main() {
}
