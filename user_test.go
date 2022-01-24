package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	testCases := []struct {
		desc     string
		User     *UserDetails
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			User: &UserDetails{
				Name:    "gopi",
				Age:     23,
				Address: "vzgnrm",
				Delete:  true,
			},
			Mock: []interface{}{
				mock.ExpectExec("INSERT INTO user").WithArgs("gopi", 23, "vzgnrm", true).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			User: &UserDetails{
				Name:    "sudheer",
				Age:     23,
				Address: "vzgnrm",
				Delete:  true,
			},
			Mock: []interface{}{
				mock.ExpectExec("INSERT INTO user").WithArgs("sudheer", 23, "vzgnrm", true).WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			err := InsertUser(db, ts.User)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}

}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUpdate := []struct {
		desc     string
		User     *UserDetails
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			User: &UserDetails{
				Name:    "gopi",
				Age:     23,
				Address: "vzgnrm",
				Delete:  true,
			},
			Mock: []interface{}{
				mock.ExpectExec("update user").WithArgs(23, "vzgnrm", true, "gopi").WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			User: &UserDetails{
				Name:    "sudheer",
				Age:     23,
				Address: "vzgnrm",
				Delete:  true,
			},
			Mock: []interface{}{
				mock.ExpectExec("update user set Age=?,Address=?,Delete=? Where Name=?;").WithArgs(23, "vzgnrm", true, "sudheer").WillReturnError(errors.New("couldn't exec query")),
			},
			expecErr: errors.New("couldn't exec query"),
		},
	}
	for _, ts := range testUpdate {
		t.Run(ts.desc, func(t *testing.T) {
			err := UpdateUser(db, ts.User)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {

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
			err := DeleteUser(db)
			if err != nil && !reflect.DeepEqual(tes.expecErr, err) {
				fmt.Print("expected ", tes.expecErr, "obtained", err)
			}
		})
	}
}

func TestReadUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testFetch := []struct {
		desc     string
		name     string
		Mock     []interface{}
		expecRes *UserDetails
		expecErr error
	}{
		{
			desc: "Success case",
			name: "gopi",
			Mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user where Name=?;").WithArgs("gopi").
					WillReturnRows(sqlmock.NewRows([]string{"name", "age", "address", "delete"}).
						AddRow("gopi", 21, "vzg", false)),
			},
			expecRes: &UserDetails{"gopi", 21, "vzg", false},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			name: "sudheer",
			Mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user where Name=?;").WithArgs("sudheer").WillReturnError(errors.New("couldn't exec query")),
			},
			expecRes: nil,
			expecErr: errors.New("couldn't exec query"),
		},
	}
	for _, ts := range testFetch {
		t.Run(ts.desc, func(t *testing.T) {
			res, err := ReadUser(db, ts.name)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
			if res != nil && !reflect.DeepEqual(ts.expecRes, res) {
				fmt.Print("expected ", ts.expecRes, "obtained", res)
			}
		})
	}
}
