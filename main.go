package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id            int
	name, address string
}

type SqlDB struct {
	db *sql.DB
}

func connectDb() (*sql.DB, error) {
	connection, error := sql.Open("mysql", "root:password@tcp(127.0.0.1)/test")
	if error != nil {
		return nil, errors.New("DATABASE CONNECTION ERROR")
	}
	return connection, nil
}

func (db *SqlDB) InsertStudent(id int, name string, address string) error {
	_, error := db.db.Exec("insert into student(id, name, address) values(?, ?, ?)", id, name, address)

	if id < 0 {
		return errors.New("STUDENT ID IS INVALID")
	}

	if error != nil {
		return errors.New("FAILED TO ADD NEW STUDENT")
	}
	return nil
}

func (db *SqlDB) DeleteStudent(id int) error {
	_, error := db.db.Exec("delete from student where id=?", id)

	if id < 0 {
		return errors.New("FAILED TO DELETE THE STUDENT")
	}

	if error != nil {
		return errors.New("STUDENT NOT EXISTS")
	}
	return nil
}

func (db *SqlDB) UpdateStudent(id int, name string, address string) error {
	_, error := db.db.Exec("update student set name=?, address=? where id=?", name, address, id)

	if id < 0 {
		return errors.New("INVALID STUDENT ID")
	}

	if error != nil {
		return errors.New("FAILED TO UPDATE STUDENT")
	}
	return nil
}

func (db *SqlDB) SelectStudentById(id int) (Student, error) {
	stud := Student{}
	con := db.db.QueryRow("select * from student where id=?", id)

	error := con.Scan(&stud.id, &stud.name, &stud.address)

	if error != nil {
		return stud, errors.New("ERROR IN SELECTING")
	}

	return stud, nil
}

func (db *SqlDB) SelectAllStudents() ([]Student, error) {
	StudentList := make([]Student, 10)
	con, error := db.db.Query("select * from student")

	if error != nil {
		return StudentList, errors.New("ERROR IN SELECTING STUDENTS")
	}

	defer con.Close()

	for con.Next() {
		stud := Student{}
		error := con.Scan(&stud.id, &stud.name, &stud.address)
		if error != nil {
			return StudentList, errors.New("ERROR IN SELECTING STUDENTS")
		}
		StudentList = append(StudentList, stud)
	}
	return StudentList, nil
}

func main() {
	connection, _ := connectDb()
	dbconn := SqlDB{connection}

	fmt.Println(dbconn.InsertStudent(4, "madhu", "Banglore"))
	fmt.Println(dbconn.SelectStudentById(4))
	fmt.Println(dbconn.UpdateStudent(4, "suhas", "pune"))
	fmt.Println(dbconn.DeleteStudent(4))
}
