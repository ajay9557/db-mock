package userDB

import (
	"database/sql"
	"errors"
	"fmt"
)

// user struct to hold the values of the table
type User struct {
	id          int
	name        string
	age         int
	deletedFlag int
}

func CreateTable(db *sql.DB, tableName string, tableSchema string) error {
	// form the query string according to the table schema
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( %s )", tableName, tableSchema)

	_, err := db.Exec(query)
	if err != nil {
		return errors.New("error creating the table")
	}
	return nil
}

func GetValuesById(db *sql.DB, tableName string, pk int) (user User, err error) {
	// form the query string according to the primary key
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)

	// execute the query and save the result in user struct variable
	err = db.QueryRow(query, pk).Scan(&user.id, &user.name, &user.age, &user.deletedFlag)
	if err != nil {
		return user, sql.ErrNoRows
	}
	return user, nil
}

func InsertValues(db *sql.DB, tableName string, id, age int, name string) (err error) {

	// form the query string according to the inputs
	query := fmt.Sprintf("INSERT INTO user(id, name, age) VALUES (?, ?, ?)")

	_, err = db.Exec(query, id, name, age)
	if err != nil {
		return errors.New("error inserting values")
	}
	return nil
}

func DeleteRecord(db *sql.DB, tableName string, pk int) error {
	_, err := db.Exec("UPDATE ? SET deletedFlag = 1 WHERE id = ?", tableName, pk)
	if err != nil {
		return errors.New("error deleting the selected record")
	}
	return nil
}

func UpdateRecord(db *sql.DB, tableName string, pk int, column string, value interface{}) error {
	_, err := db.Exec("UPDATE ? SET ? = ? WHERE id = ?", tableName, column, value, pk)
	if err != nil {
		return errors.New("error updating the records")
	}
	return nil
}
