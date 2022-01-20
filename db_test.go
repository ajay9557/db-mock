package main

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestReadById(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Naruto", 21, "Surat")

	tests := []struct {
		desc      string
		id        int
		expected  User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expected: User{id: 1, name: "Naruto", age: 21, address: "Surat"}, id: 1, mockQuery: mock.ExpectQuery("select id from users where id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", expected: User{}, id: 2, mockQuery: nil},
		{desc: "Case3", expected: User{}, id: 1, mockQuery: mock.ExpectQuery("select id from users").WillReturnError(errors.New("Connection lost"))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u, err := ReadById(db, test.id)

			if err != nil && !errors.Is(err, sql.ErrNoRows) && !reflect.DeepEqual(u, test.expected) {
				t.Errorf("Expected: %v, Got: %v", rows, u)
			}
		})
	}
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Println(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Naruto", 21, "Japan").AddRow(2, "Ichigo", 18, "America")
	rows2 := sqlmock.NewRows([]string{"id", "name", "age", "address"})

	tests := []struct {
		desc      string
		expected  []User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expected: []User{
			{id: 1, name: "Naruto", age: 21, address: "Japan"},
			{id: 2, name: "Ichigo", age: 18, address: "America"},
		}, mockQuery: mock.ExpectQuery("select id from users").WillReturnRows(rows)},
		{desc: "Case2", expected: []User{}, mockQuery: mock.ExpectQuery("select id from users").WillReturnRows(rows2)},
		{desc: "Case3", expected: []User{}, mockQuery: mock.ExpectQuery("select id from users").WillReturnError(errors.New("Connection lost"))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			users, err := Read(db)

			if !reflect.DeepEqual(users, test.expected) && err != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, users)
			}
		})
	}
}

func TestInsertRecord(t *testing.T) {
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
		{desc: "Case1", expected: 1, mockCall: mock.ExpectExec("insert into users(name, age, address) values(?, ?, ?)").WithArgs("Ridhdhish", 21, "Surat").WillReturnResult(sqlmock.NewResult(1, 1))},
		{desc: "Case2", expected: 0, mockCall: mock.ExpectExec("insert into users(name, age, address) values(?, ?, ?)").WithArgs("Ridhdhish", 21).WillReturnResult(sqlmock.NewResult(0, 0))},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, _ := InsertRecord(db, 21, "Ridhdhish", "Surat")

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

	mock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Ridhdhish", 21, "Surat")

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

	mock.NewRows([]string{"id", "name", "age", "address"}).AddRow(2, "Ridhdhish", 21, "Surat")

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertedId, _ := DeleteRecord(db, test.id)

			if lastInsertedId != test.expected {
				t.Errorf("Expected: %d, Got: %d", 1, lastInsertedId)
			}
		})
	}
}
