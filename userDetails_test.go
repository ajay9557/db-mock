package main

import (
	"errors"
	//"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcsInsert := []struct {
		desc string
		Id   int
		Age  int
		Name string
		Add  string
		Mock []interface{}
		err  error
	}{
		{
			desc: "Success",
			Id:   1,
			Age:  12,
			Name: "test",
			Add:  "beng",
			//the below one is for having multile methods to test like insert,update,delete.
			Mock: []interface{}{
				mock.ExpectExec(`insert into test2 values`).WithArgs(1, "test", 12, "beng").WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc: "Failure",
			Id:   2,
			Age:  34,
			Name: "test1",
			Add:  "hyd",
			Mock: []interface{}{
				mock.ExpectExec("insert into test2 values").WillReturnError(errors.New("t")),
			},
			err: errors.New("t"), //if we write error message different here it gives error.
		},
	}
	for _, tc := range tcsInsert {
		err := addDetails(db, tc.Id, tc.Name, tc.Age, tc.Add)
		if tc.Id < 1 {
			return
		}
		t.Run(tc.desc, func(t *testing.T) {
			// if tc.err == nil && err != nil {
			// 	t.Errorf("There is an error : %s", err)
			// }
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcsUpdate := []struct {
		desc string
		Id   int
		Age  int
		Name string
		Add  string
		Mock []interface{}
		err  error
	}{
		{
			desc: "Success",
			Id:   1,
			Age:  12,
			Name: "test",
			Add:  "beng",
			//the below one is for having multile methods to test like insert,update,delete.
			Mock: []interface{}{
				mock.ExpectExec(`update test2`).WithArgs(12, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc: "Failure",
			Id:   2,
			Age:  34,
			Name: "test1",
			Add:  "hyd",
			Mock: []interface{}{
				mock.ExpectExec(`update test2`).WithArgs(34, 2).WillReturnError(errors.New("t")),
			},
			err: errors.New("t"), //if we write error message different here it gives error.
		},
	}
	for _, tc := range tcsUpdate {
		err := updateDetails(db, tc.Age, tc.Id)
		if tc.Id < 1 {
			return
		}
		t.Run(tc.desc, func(t *testing.T) {
			// if tc.err == nil && err != nil {
			// 	t.Errorf("There is an error : %s", err)
			// }
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcsDelete := []struct {
		desc string
		Id   int
		Age  int
		Name string
		Add  string
		Mock []interface{}
		err  error
	}{
		{
			desc: "Success",
			Id:   1,
			Age:  12,
			Name: "test",
			Add:  "beng",
			//the below one is for having multile methods to test like insert,update,delete.
			Mock: []interface{}{
				mock.ExpectExec(`delete from test2`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc: "Failure",
			Id:   2,
			Age:  34,
			Name: "test1",
			Add:  "hyd",
			Mock: []interface{}{
				mock.ExpectExec(`delete from test2`).WithArgs(2).WillReturnError(errors.New("t")),
			},
			err: errors.New("t"), //if we write error message different here it gives error.
		},
	}
	for _, tc := range tcsDelete {
		err := deleteDetails(db, tc.Id)
		if tc.Id < 1 {
			return
		}
		t.Run(tc.desc, func(t *testing.T) {
			// if tc.err == nil && err != nil {
			// 	t.Errorf("There is an error : %s", err)
			// }
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}

func Test_Search(t *testing.T) {
	query := "SELECT id,name,age,adress FROM test2"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}
	tcsSearch := []struct {
		desc     string
		Id       int
		expected Tag
		Mock     []interface{}
	}{
		{
			desc:     "Success",
			Id:       2,
			expected: Tag{2, "Sudheer", 22, "Hyderabad"},
			Mock:     []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "Name", "Age", "adress"}).AddRow(2, "Sudheer", 22, "Hyderabad"))},
		},
	}
	for _, tc := range tcsSearch {
		res, err := searchDetails(db)
		if err != nil {
			return
		}
		t.Run(tc.desc, func(t *testing.T) {
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}
