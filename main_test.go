package main

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// _ "github.com/DATA-DOG/go-sqlmock"
	// _ "github.com/stretchr/testify/assert"
)

// var u = []users{{
// 	Id:      1,
// 	Name:    "Anu",
// 	Age:     51,
// 	Address: "repalle",
// },
// 	{
// 		Id:      2,
// 		Name:    "Veedaa",
// 		Age:     62,
// 		Address: "vzd",
// 	},
// }

// func TestReadRecordsById(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	query := "select id,name,age,address from UUser where id = ?"
// 	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(u[0].Id, u[0].Name, u[0].Age, u[0].Address)
// 	mock.ExpectQuery(query).WithArgs(u[0].Id).WillReturnRows(rows)
// 	// _, err = ReadRecordsById(db, int(u[0].Id))
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// }

func TestReadRecordsById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Naruto", 21, "Surat")

	tests := []struct {
		desc      string
		id        int
		expected  users
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expected: users{Id: 1, Name: "Naruto", Age: 21, Address: "Surat"}, id: 1, mockQuery: mock.ExpectQuery("select * from users where id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", expected: users{}, id: 2, mockQuery: nil},
		{desc: "Case3", expected: users{}, id: 1, mockQuery: mock.ExpectQuery("select id from users").WillReturnError(errors.New("Connection lost"))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u, err := ReadRecordsById(db, test.id)

			if err != nil && !errors.Is(err, sql.ErrNoRows) && !reflect.DeepEqual(u, test.expected) {
				t.Errorf("Expected: %v, Got: %v", rows, u)
			}
		})
	}
}

func TestReadRecords(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Naruto", 21, "Japan").AddRow(2, "Ichigo", 18, "America")
	rows2 := sqlmock.NewRows([]string{"id", "name", "age", "address"})

	tests := []struct {
		desc      string
		expected  []users
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expected: []users{
			{Id: 1, Name: "Naruto", Age: 21, Address: "Japan"},
			{Id: 2, Name: "Ichigo", Age: 18, Address: "America"},
		}, mockQuery: mock.ExpectQuery("select * from users").WillReturnRows(rows)},
		{desc: "Case2", expected: []users{}, mockQuery: mock.ExpectQuery("select * from users").WillReturnRows(rows2)},
		{desc: "Case3", expected: []users{}, mockQuery: mock.ExpectQuery("select * from users").WillReturnError(errors.New("Connection lost"))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			users, err := ReadRecords(db)

			if !reflect.DeepEqual(users, test.expected) && err != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, users)
			}
		})
	}
}

func TestCreateRecord(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	tests := []struct {
		desc     string
		expected int64
		mockCall *sqlmock.ExpectedExec
	}{
		{desc: "Case1", expected: 1, mockCall: mock.ExpectExec("insert into users(name, age, address) values(?, ?, ?)").WithArgs("Anu", 21, "Repaale").WillReturnResult(sqlmock.NewResult(1, 1))},
		{desc: "Case2", expected: 0, mockCall: mock.ExpectExec("insert into users(name, age, address) values(?, ?, ?)").WithArgs("Anu", 21).WillReturnResult(sqlmock.NewResult(0, 0))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, _ := CreateRecord(db, "anusri", 22, "Repalle")

			if user != test.expected && err != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	tests := []struct {
		desc     string
		expected int64
		mockCall *sqlmock.ExpectedExec
	}{
		{desc: "Case1", expected: 1, mockCall: mock.ExpectExec("update users set name = ? where id = ?").WithArgs("Naruto", 1).WillReturnResult(sqlmock.NewResult(1, 1))},
		{desc: "Case2", expected: 0, mockCall: mock.ExpectExec("update users set name = ? where id = ?").WithArgs("Naruto", 2).WillReturnResult(sqlmock.NewResult(0, 0))},
	}

	mock.NewRows([]string{"id", "name", "age", "address"}).AddRow(12, "Anusri", 21, "Rpl")

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := UpdateRecord(db, "Naruto", 1)

			if affectedRows != test.expected && err != nil {
				t.Errorf("Expected: %d, Got: %d", 1, affectedRows)
			}
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	tests := []struct {
		desc     string
		id       int
		expected int64
		mockCall *sqlmock.ExpectedExec
	}{
		{desc: "Case1", id: 2, expected: 2, mockCall: mock.ExpectExec("delete from users where id = ?").WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))},
		{desc: "Case2", id: 8, expected: 0, mockCall: mock.ExpectExec("delete from users where id = ?").WithArgs(8).WillReturnResult(sqlmock.NewResult(0, 0))},
	}

	mock.NewRows([]string{"id", "name", "age", "address"}).AddRow(2, "Anusri", 21, "Rpl")

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertedId, _ := DeleteRecord(db, test.id)

			if lastInsertedId != test.expected {
				t.Errorf("Expected: %d, Got: %d", 1, lastInsertedId)
			}
		})
	}
}
