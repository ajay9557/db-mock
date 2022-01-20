package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/DATA-DOG/go-sqlmock"
)

type user struct {
	id      int
	name    string
	age     int
	address string
	delete  bool
}

func CreateTable(db *sql.DB) {
	query := "create table user(id int not null,name varchar(50) not null,age int,address varchar(100))"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println("error while creating table", err)
		return
	}
	fmt.Println("table created successfully", res)
}

func Insert(db *sql.DB, usr *user) error {
	query := "insert into user values(?,?,?,?)"
	stmt, err := db.Prepare(query)
	defer func() {
		if err == nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("error in preparing statement in insert", err)
		return err
	}
	result, err := stmt.Exec(usr.id, usr.name, usr.age, usr.address)
	if err != nil {
		fmt.Println("error while executing insert statement", err)
		return err
	}
	id, _ := result.LastInsertId()
	fmt.Println("last inserted id: ", id)
	return nil
}

func Delete(db *sql.DB, id int) error {
	query := "delete from user where id = (?)"
	result, err := db.Exec(query, id)
	if err != nil {
		//fmt.Println("error while executing delete statement")
		return err
	}
	rows, _ := result.RowsAffected()
	//fmt.Println("rows affected: ", rows)
	if rows != 1 {
		return errors.New("no id found to delete")
	}
	return nil
}

func ShowAll(db *sql.DB) error {
	query := "select * from user"
	rows, err := db.Query(query)
	defer func() {
		if err == nil {
			rows.Close()
		}
	}()

	//if err == sql.ErrNoRows {
	//	return err
	//}
	if err != nil {
		//fmt.Println("error in executing select statement in showall", err)
		return err
	}
	//} else {
	//	defer rows.Close()
	//}
	for rows.Next() {
		usr := user{}
		rows.Scan(&usr.id, &usr.name, &usr.age, &usr.address)
		fmt.Println(usr)
	}
	return nil
}

func Update(db *sql.DB, usr *user) error {
	query := "update user set name = ? where id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error while preparing update stmt")
		return err
	}
	result, err := stmt.Exec(usr.name, usr.id)

	if err != nil {
		fmt.Println("error while executing update stmt")
		return err
	}
	rA, _ := result.RowsAffected()
	fmt.Println("rows Affected : ", rA)
	if rA != 1 {
		return errors.New("no id found to update")
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", "test:1234@/test")
	defer db.Close()
	if err != nil {
		fmt.Println("error while opening sql")
	}
	CreateTable(db)
	usr := user{3, "himanshu", 25, "hajipur", false}
	//usr1 := user{1, "rahul", 17, "hajipur", false}

	Insert(db, &usr)
	//Insert(db, &usr1)
	Delete(db, 1)
	//Update(db, &usr1)
	ShowAll(db)

}
