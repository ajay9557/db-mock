package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
)

var u = user{
	id:      1,
	name:    "himanshu",
	age:     25,
	address: "hajipur",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func Test_Insert(t *testing.T) {
	db, mock := NewMock()
	query := "insert into user values(?,?,?,?)"
	//query1 := "insert into user values(?,?,?,?)"

	testcases := []struct {
		usr  user
		err  error
		mock []interface{}
	}{
		{
			usr: u,
			err: nil,
			mock: []interface{}{
				mock.ExpectPrepare(query).ExpectExec().WithArgs(u.id, u.name, u.age, u.address).WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
		{
			usr: u,
			err: sql.ErrConnDone,
			mock: []interface{}{
				mock.ExpectPrepare(query).WillReturnError(sql.ErrConnDone),
			},
		},
		{
			usr: u,
			err: sql.ErrConnDone,
			mock: []interface{}{
				mock.ExpectPrepare(query).ExpectExec().WithArgs(u.id, u.name, u.age, u.address).WillReturnError(sql.ErrConnDone),
			},
		},
	}

	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(u.id, u.name, u.age, u.address).WillReturnResult(sqlmock.NewResult(0, 1))
	for _, tcs := range testcases {
		err := Insert(db, &tcs.usr)
		if !reflect.DeepEqual(err, tcs.err) {
			t.Error("error in test insert")
		}
	}
}

func Test_Delete(t *testing.T) {
	db, mock := NewMock()
	query := "delete from user where id = (?)"

	testcases := []struct {
		desc string
		id   int
		err  error
		mock []interface{}
	}{
		{
			"testcase-1",
			1,
			nil,
			[]interface{}{
				mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
		{
			"testcase-2",
			6,
			errors.New("no id found to delete"),
			[]interface{}{
				mock.ExpectExec(query).WithArgs(6).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
		{
			"testcase-3",
			6,
			sql.ErrConnDone,
			[]interface{}{
				mock.ExpectExec(query).WithArgs(6).WillReturnError(sql.ErrConnDone),
			},
		},
	}
	for _, tcs := range testcases {
		err := Delete(db, tcs.id)

		if !reflect.DeepEqual(err, tcs.err) {
			t.Error("error : ", tcs.desc, err)
		}
	}
}

func Test_ShowAll(t *testing.T) {
	db, mock := NewMock()
	query := "select * from user"
	testcases := []struct {
		desc string
		err  error
		mock []interface{}
	}{
		{
			"test-1",
			sql.ErrNoRows,
			[]interface{}{
				mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows),
			},
		},
		{
			"test-2",
			nil,
			[]interface{}{
				mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "address"}).AddRow(1, "himanshu", 25, "hajipur")),
			},
		},
	}

	for _, tcs := range testcases {
		err := ShowAll(db)
		if !reflect.DeepEqual(err, tcs.err) {
			t.Errorf("expected %v, got %v", tcs.err, err)
		}
	}
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	query := "update user set name = ? where id = ?"

	testcases := []struct {
		usr  user
		err  error
		mock []interface{}
	}{
		{
			u,
			nil,
			[]interface{}{
				mock.ExpectPrepare(query).ExpectExec().WithArgs(u.name, u.id).WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
		{
			u,
			sql.ErrConnDone,
			[]interface{}{
				mock.ExpectPrepare(query).WillReturnError(sql.ErrConnDone),
			},
		},
		{
			u,
			sql.ErrConnDone,
			[]interface{}{
				mock.ExpectPrepare(query).ExpectExec().WithArgs(u.name, u.id).WillReturnError(sql.ErrConnDone),
			},
		},
		{
			u,
			errors.New("no id found to update"),
			[]interface{}{
				mock.ExpectPrepare(query).ExpectExec().WithArgs(u.name, u.id).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
	}

	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(u.name, u.id).WillReturnResult(sqlmock.NewResult(0, 1))
	for _, tcs := range testcases {
		err := Update(db, &tcs.usr)
		if !reflect.DeepEqual(err, tcs.err) {
			fmt.Println("error in test insert")
		}
	}
}
