package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertTable(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	testCases := []struct {
		desc     string
		name     string
		age      int
		address  string
		delete   bool
		Mock     []interface{}
		expecErr error
	}{
		{
			desc:    "Success case",
			name:    "gopi",
			age:     21,
			address: "vzg",
			delete:  false,
			Mock: []interface{}{
				mock.ExpectExec("INSERT INTO user").WithArgs("gopi", 21, "vzg", false).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc:    "Failure case",
			name:    "sudheer",
			age:     23,
			address: "vzrt",
			delete:  true,
			Mock: []interface{}{
				mock.ExpectExec("INSERT INTO user").WithArgs("sudheer", 23, "vzrt", true).WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			err := InsertTable(db, ts.name, ts.age, ts.address, ts.delete)
			if ts.expecErr == nil && err != nil {
				t.Errorf("error occured %s", err)
			}
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}

}

func TestUpdateDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUpdate := []struct {
		desc     string
		name     string
		age      int
		address  string
		delete   bool
		Mock     []interface{}
		expecErr error
	}{
		{
			desc:    "Success case",
			name:    "gopi",
			age:     23,
			address: "vzgnrm",
			delete:  true,
			Mock: []interface{}{
				mock.ExpectExec("update user").WithArgs(23, "vzgnrm", true, "gopi").WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc:    "Failure case",
			name:    "sudheer",
			age:     23,
			address: "vzgnrm",
			delete:  true,
			Mock: []interface{}{
				mock.ExpectExec("update user").WithArgs(23, "vzgnrm", true, "gopi").WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}
	for _, ts := range testUpdate {
		t.Run(ts.desc, func(t *testing.T) {
			err := UpdateDB(db, ts.age, ts.address, ts.delete, ts.name)
			if ts.expecErr == nil && err != nil {
				t.Errorf("error occured %s", err)
			}
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}
}

func TestDeleteDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	testCases := []struct {
		desc     string
		delete   bool
		Mock     []interface{}
		expecErr error
	}{
		{
			desc:   "Success case",
			delete: true,
			Mock: []interface{}{
				mock.ExpectExec("DELETE From user ").WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc:   "Failure case",
			delete: false,
			Mock: []interface{}{
				mock.ExpectExec("DELETE From user ").WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}

	for _, tes := range testCases {

		t.Run(tes.desc, func(t *testing.T) {
			err := DeleteDB(db)
			if tes.expecErr == nil && err != nil {
				t.Errorf("error occured %s", err)
			}
			if err != nil && !reflect.DeepEqual(tes.expecErr, err) {
				fmt.Print("expected ", tes.expecErr, "obtained", err)
			}
		})
	}
}

func TestFetchTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testFetch := []struct {
		desc     string
		name     string
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			name: "gopi",
			Mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user").WithArgs("gopi").
					WillReturnRows(sqlmock.NewRows([]string{"name", "age", "address", "delete"}).
						AddRow("gopi", 21, "vzg", false)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			name: "sudheer",
			Mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user").WithArgs("sudheer").WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}
	for _, ts := range testFetch {
		t.Run(ts.desc, func(t *testing.T) {
			res, err := FetchTable(db, ts.name)
			if err != nil && ts.expecErr == nil {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
			fmt.Println(res)
		})
	}
}
