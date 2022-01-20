package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetchRecords(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Naruto", 21, "Surat")

	// mock.ExpectBegin()

	tests := []struct {
		desc      string
		id        int
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, mockQuery: mock.ExpectQuery("select id from users where id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", id: 2, mockQuery: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u, err := FetchRecords(db, test.id)

			if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
		expected  int64
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expected: 2, mockQuery: mock.ExpectQuery("select id from users").WillReturnRows(rows)},
		{desc: "Case2", expected: 0, mockQuery: mock.ExpectQuery("select id from users").WillReturnRows(rows2)},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, err := Read(db)

			if err != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, affectedRows)
			}

			if affectedRows != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, affectedRows)
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
			affectedRows, _ := InsertRecord(db, 21, "Ridhdhish", "Surat")

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", 1, affectedRows)
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

			if affectedRows != test.expected {
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
		{desc: "Case1", id: 1, expected: 1, mockCall: mock.ExpectExec("delete from users where id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))},
		{desc: "Case2", id: 8, expected: 0, mockCall: mock.ExpectExec("delete from users where id = ?").WithArgs(8).WillReturnResult(sqlmock.NewResult(0, 0))},
	}

	mock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "Ridhdhish", 21, "Surat")

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := DeleteRecord(db, test.id)

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", 1, affectedRows)
			}
		})
	}
}
