package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserDetails struct {
	Name    string
	Age     int
	Address string
	Delete  bool
}

func UpdateDB(db *sql.DB, age int, address string, del bool, name string) (err error) {
	updateQ := "update user set Age=?, Address=?, Del=? where Name = ?;" //Update query
	_, err = db.Exec(updateQ, age, address, del, name)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func FetchTable(db *sql.DB, name string) (user UserDetails, err error) {
	ReadQ := "SELECT * FROM user where name=?;" //Fetch values from table
	user = UserDetails{}
	rows, err := db.Query(ReadQ, name)
	if err != nil {
		return user, errors.New("couldn't exec query")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Age, &user.Address, &user.Delete)
		if err != nil {
			return user, errors.New("couldn't exec query")
		}
	}
	return user, nil
}

func InsertTable(db *sql.DB, Name string, Age int, Address string, Del bool) (err error) {
	// Create and insert values into table
	insertQ := "INSERT INTO user(Name, Age, Address,Del) VALUES(?,?,?,?);" //Insert query
	_, err = db.Exec(insertQ, Name, Age, Address, Del)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func DeleteDB(db *sql.DB) (err error) {
	deleteQ := "DELETE From user where Del=?;"
	_, err = db.Exec(deleteQ, true)
	if err != nil {
		return errors.New("couldn't exec query")
	}
	return nil
}

func main() {
	//Connecting to database
	db, err := sql.Open("mysql", "gopi:gopi@123@tcp(0.0.0.0:3306)/userDetails")
	if err != nil {
		fmt.Printf(err.Error())
	}

	defer db.Close()
	users := []UserDetails{{Name: "Gopi", Age: 21, Address: "vzg", Delete: false},
		{Name: "Sudheer", Age: 21, Address: "Nzm", Delete: false},
		{Name: "Ram", Age: 21, Address: "Vjw", Delete: false},
		{Name: "Dummy", Age: 21, Address: "***", Delete: true},
	}
	InsertTable(db, users[0].Name, users[0].Age, users[0].Address, users[0].Delete)

	upDetails := []UserDetails{{Name: "Gopi", Age: 23, Address: "vskp", Delete: false},
		{Name: "Sudheer", Age: 20, Address: "Nzmbd", Delete: false},
		{Name: "Ram", Age: 20, Address: "Vjd", Delete: false}}
	UpdateDB(db, upDetails[0].Age, upDetails[0].Address, upDetails[0].Delete, upDetails[0].Name)

}
