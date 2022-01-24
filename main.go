package main

import (
	"database/sql"
	"db-mock/userDB"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "vips:1234@tcp(0.0.0.0:3306)/test")
	if err != nil {
		log.Printf("error connection to database : Error %v", err)
	}

	defer db.Close()

	// Create the table
	err = userDB.CreateTable(db, "user", "id INT NOT NULL,name VARCHAR(255) NOT NULL,age INT NOT NULL,deleted INT NOT NULL DEFAULT 0,PRIMARY KEY (id)")
	if err != nil {
		log.Printf("error creating table : Error %v", err)
	}

	// Insert the data
	names := []string{
		"Vipul",
		"Jane",
		"John",
		"Hello",
		"World",
	}
	for i := 0; i < 5; i++ {
		err = userDB.InsertValues(db, "user", i, rand.Intn(40), names[i])
		if err != nil {
			log.Printf("error inserting data in the table : Error %v", err)
		}
	}

	// Get the data according to the id
	user, err := userDB.GetValuesById(db, "user", 1)
	if err != nil {
		log.Printf("error getting values : Error %v", err)
	}
	fmt.Println(user)

	// Delete a record by setting the deleted flag to 1
	err = userDB.DeleteRecord(db, "user", 1)
	if err != nil {
		log.Printf("error deleting row : Error %v", err)
	}

	// Update the record
	err = userDB.UpdateRecord(db, "user", 1, "name", "Heya")
	if err != nil {
		log.Printf("error updating row : Error %v", err)
	}

	_, err = db.Exec("DROP TABLE user")
	if err != nil {
		log.Printf("error executing query")
	}
}
