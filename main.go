package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID     int
	Name   string
	Age    int
	adress string
}

func main() {
	db, err := sql.Open("mysql", "SudheerKumar:Puppala@tcp(127.0.0.1:3306)/Practice")
	if err != nil {
		fmt.Print(err)
	}

	// tag := []Tag{
	// 	{1, "Sudheer", 22, "Hyderabad"},
	// 	{2, "Puppala", 20, "Chennai"},
	// 	{3, "Maheshbabu", 42, "Hyderabad"},
	// 	{4, "Ram", 32, "Chennai"},
	// 	{5, "Yash", 19, "Bangalore"},
	// 	{6, "Sam", 18, "Kerala"},
	// 	{7, "Charan", 21, "Chennai"},
	// 	{8, "Jon", 21, "Chennai"},
	// 	{9, "Jash", 25, "Hyderabad"},
	// }

	addDetails(db, 4, "Ram", 32, "Chennai")
	updateDetails(db, 35, 4)
	deleteDetails(db, 3)
	searchDetails(db)

}

func addDetails(db *sql.DB, ID int, Name string, Age int, adress string) (err error) {
	if _, err = db.Exec(`insert into test2 values(?,?,?,?)`, ID, Name, Age, adress); err != nil {
		return err
	}
	return nil
}

func updateDetails(db *sql.DB, Age, ID int) (err error) {
	if _, err = db.Exec("update test2 set age = ? where id = ?", Age, ID); err != nil {
		return
	}
	return nil
}

func deleteDetails(db *sql.DB, ID int) (err error) {
	if _, err = db.Exec("delete from test2 where id=?", ID); err != nil {
		return
	}
	return nil
}
func searchDetails(db *sql.DB) (t Tag, err error) {
	t = Tag{}
	results, err := db.Query("SELECT id,name,age,adress FROM test2")
	if err != nil {
		fmt.Print(err)
	}
	for results.Next() {
		err = results.Scan(&t.ID, &t.Name, &t.Age, &t.adress)
		if err != nil {
			return t, err
			//fmt.Print(err)
		}
	}
	return t, nil
}
