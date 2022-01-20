package userDB

import (
	"database/sql"
	"log"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTable(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"
	tableSchema := `id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	age INT NOT NULL,
	deleted INT NOT NULL DEFAULT 0,
	PRIMARY KEY (id)`

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing the db connection")
		}
	}(db)

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS user").WillReturnResult(sqlmock.NewResult(1, 1))

	if err := CreateTable(db, tableName, tableSchema); err != nil {
		log.Printf("error returned by CreateTable: %s", err)
	}
}

func TestGetValuesById(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing the db connection")
		}
	}(db)

	rows := sqlmock.NewRows([]string{"id", "name", "age", "deleted"}).
		AddRow(1, "John", 20, 0)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, age, deleted FROM user WHERE id = ?")).
		WithArgs(1).
		WillReturnRows(rows)

	user, err := GetValuesById(db, tableName, 1)
	if err != nil {
		log.Printf("error returned by GetValuesById: %s", err)
	}
	if !reflect.DeepEqual(user, User{1, "John", 20, 0}) {
		t.Errorf("GetValuesById() = %v, want %v", user, User{1, "John", 20, 0})
	}
}

func TestInsertValues(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing the db connection")
		}
	}(db)

	mock.ExpectExec("INSERT INTO user").
		WithArgs(1, "John", 20, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = InsertValues(db, tableName, 1, "John", 20, 0)
	//if !reflect.DeepEqual(user, User{1, "John", 20, 0}) {
	//	t.Errorf("InsertValues() = %v, want %v", user, User{1, "John", 20, 0})
	//}
	if err != nil {
		log.Printf("error returned by InsertValues: %s", err)
	}
}

func TestDeleteRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing the db connection")
		}
	}(db)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE ? SET deletedFlag = 1 WHERE id = ?")).
		WithArgs(tableName, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteRecord(db, "user", 1); err != nil {
		log.Printf("error returned by DeleteRecord: %s", err)
	}
}

func TestUpdateRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing the db connection")
		}
	}(db)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE ? SET ")).
		WithArgs(tableName, "name", "Jane", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = UpdateRecord(db, tableName, 1, "name", "Jane")
	if err != nil {
		log.Printf("an error occured while updating %v", err)
	}
}
