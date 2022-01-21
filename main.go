package main

import (
	"database/sql"
	"errors"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/gorilla/mux"
	// _ "github.com/jinzhu/gorm"
)

type users struct {
	Id      int64
	Name    string
	Age     int
	Address string
}

var aryofusers []users

func main() {
	db, err := sql.Open("mysql", "root:satyasusi@tcp(127.0.0.1:3306)/test")
	CheckNilError(err)
	defer db.Close()
	_, err = db.Exec("create table if not exists UUser(id int primary key AUTO_INCREMENT,name varchar(255) not null,age int not null,address varchar(255) not null)")
	CheckNilError(err)
	// ReadRecords(db)
	// PrintUsers(aryofusers)
	CreateRecord(db, "anu", 21, "repalle")
	CreateRecord(db, "veda", 62, "VZd")
	// PrintUsers(aryofusers)
	UpdateRecord(db, "anusri", 29)
	aa, err := ReadRecordsById(db, 28)
	if err != nil {
		fmt.Println(err)
	}
	// PrintUsers(aa)
	fmt.Println(aa)
	rows, _ := ReadRecords(db)
	fmt.Println("All Records: ", rows)

	// PrintUsers(aryofusers)
	DeleteRecord(db, 31)
	// PrintUsers(aryofusers)

}

func CreateRecord(db *sql.DB, name string, age int, address string) (int64, error) {

	// Creating an user into UUser db
	result, err := db.Exec("Insert into UUser(name,age,address) values (?,?,?)", name, age, address)

	// CheckNilError(err)
	if err != nil {
		return 0, errors.New("couldn't insert record")
	}
	iid, _ := result.LastInsertId()
	// CheckNilError(err)
	// user.Id = iid
	// aryofusers = append(aryofusers, user)
	return iid, nil

}

func CreateRecord1(db *sql.DB, name string, age int, address string) {
	var user users
	user.Name = name
	user.Age = age
	user.Address = address
	// Creating an user into UUser db
	result, err := db.Exec("Insert into UUser(name,age,address) values (?,?,?)", name, age, address)

	CheckNilError(err)
	iid, err := result.LastInsertId()
	CheckNilError(err)
	user.Id = iid
	aryofusers = append(aryofusers, user)

}

func ReadRecords(db *sql.DB) ([]users, error) {
	rows, err := db.Query("select * from UUser")

	// CheckNilError(err)
	if err != nil {
		return []users{}, errors.New("couldn't fetch")
	}
	for rows.Next() {
		var user users
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)
		// CheckNilError(err)
		if err != nil {
			return []users{}, err
		}
		aryofusers = append(aryofusers, user)

	}
	return aryofusers, nil

}

func ReadRecords1(db *sql.DB) {
	rows, err := db.Query("select * from UUser")

	CheckNilError(err)
	// if err != nil {
	// 	return []users{}, errors.New("couldn't fetch")
	// }
	for rows.Next() {
		var user users
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)
		CheckNilError(err)
		// if err != nil {
		// 	return []users{}, err
		// }
		aryofusers = append(aryofusers, user)

	}
	// return aryofusers

}

func ReadRecordsById(db *sql.DB, id int) (users, error) {

	rows := db.QueryRow("select id,name,age,address from UUser where id = ?", id)
	if rows.Err() != nil {
		return users{}, errors.New("couldnt read")

	}

	var arrofid users

	var user users

	rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)

	arrofid = user

	return arrofid, nil

}

func ReadRecordsById1(db *sql.DB, id int) users {
	rows, err := db.Query("select id,name,age,address from UUser where id = ?", id)
	CheckNilError(err)

	var arrofid []users

	for rows.Next() {
		var user users
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address)
		CheckNilError(err)

		arrofid = append(arrofid, user)

	}

	return arrofid[0]

}

func PrintUsers(arr []users) {
	for _, v := range arr {
		// fmt.Println()
		fmt.Printf(" %v\n", v)
	}

}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func UpdateRecord(db *sql.DB, name string, id int) (int64, error) {
	result, err := db.Exec("update UUser set name = ? where id = ?", name, id)
	// CheckNilError(err)
	if err != nil {
		return 0, errors.New("couldn't update record")
	}

	noofreff, _ := result.RowsAffected()
	// CheckNilError(err)
	// fmt.Print(reff)
	return noofreff, nil

}

func UpdateRecord1(db *sql.DB, id int) {
	result, err := db.Exec("update UUser set age = age+2 where id = ?", id)
	CheckNilError(err)
	reff, err := result.RowsAffected()
	CheckNilError(err)
	fmt.Print(reff)

	for k, v := range aryofusers {
		if v.Id == int64(id) {
			// fmt.Printf("Hi : %v", v)
			v.Age = v.Age + 2
			aryofusers[k] = v
		}
	}

}

func DeleteRecord(db *sql.DB, d int) (int64, error) {
	result, err := db.Exec("delete from UUser where id = ?", d)
	// CheckNilError(err)
	if err != nil {
		return 0, errors.New("couldn't delete record")
	}
	reff, _ := result.LastInsertId()
	// CheckNilError(err)
	// fmt.Print(reff)
	return reff, nil

}

func DeleteRecord1(db *sql.DB, d int) {
	result, err := db.Exec("delete from UUser where id = ?", d)
	CheckNilError(err)
	reff, err := result.LastInsertId()
	CheckNilError(err)
	fmt.Print(reff)

	for k, v := range aryofusers {
		if v.Id == int64(d) {
			// fmt.Printf("Hi : %v", v)
			aryofusers = append(aryofusers[:k], aryofusers[k+1:]...)
		}

	}

}
