package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
	Del     bool
}

type DbUser struct {
	db *sql.DB
}

func (u *DbUser) Create(value User) (int64, int64, error) {

	query := "INSERT INTO Users(Name, Age , Address ,Del ) values(?, ?, ?, ?)"
	res, err := u.db.Exec(query, value.Name, value.Age, value.Address, value.Del)

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return 0, -1, err
	}

	affect, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return 0, -1, err
	}

	lastInsertId, err := res.LastInsertId()
	fmt.Println("Records affected", affect)

	if err != nil {

		return 0, -1, errors.New("Error in Query")
	}

	return lastInsertId, affect, nil

}

func (u *DbUser) ReadByID(id int) (*User, error) {

	user := &User{}
	query := "Select * from Users where Id = ?"
	rows, err := u.db.Query(query, id)

	if err != nil {

		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		user = &User{}
		rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Del)
	}
	return user, nil

}

func (u *DbUser) Read() (*User, error) {

	user := &User{}
	query := "Select * from Users"
	rows, err := u.db.Query(query)

	if err != nil {

		return nil, errors.New("Error in Query")
	}

	defer rows.Close()
	for rows.Next() {

		user = &User{}
		rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Del)
	}
	return user, nil

}

func (u *DbUser) Update(value string, id int) (int64, int64, error) {

	query := "Update Users Set Name = ? where Id = ?"
	result, err := u.db.Exec(query, value, id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, -1, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, -1, err
	}

	lastInsertId, err := result.LastInsertId()
	fmt.Println("Records affected", affect)

	if err != nil {
		return 0, -1, errors.New("Error in Query")
	}

	return lastInsertId, affect, nil

}

func (u *DbUser) Delete(id int) (int64, int64, error) {

	query := "DELETE FROM Users WHERE Id = ?"

	result, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, -1, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, -1, err
	}

	lastInsertId, err := result.LastInsertId()
	fmt.Println("Records affected", affect)

	if err != nil {

		return 0, -1, errors.New("Error in Query")
	}

	return lastInsertId, affect, nil
}

func main() {

	v, err := sql.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/test")
	defer v.Close()

	var u DbUser
	u.db = v

	err = v.Ping()
	if err != nil {
		log.Println("Error in accessing the connection, err: ", err)
	}

	u.Create(User{0, "Mohit", 29, "Whitefield, Bangalore", false})
	u.Read()
	u.Read()
	u.ReadByID(1)
	u.Update("Shiv", 5)
	u.Delete(11)
}
