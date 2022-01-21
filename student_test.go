package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSelectStudentById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dbHandler := &SqlDB{db}
	if err != nil {
		t.Error(err)
	}
	query := "select * from student where id=?"

	testCasesForGetStudent := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			input:    4,
			expected: Student{4, "madhu", "Banglore"},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address"}).AddRow(4, "madhu", "Banglore"))},
		},
		{
			desc:     "Case 2",
			input:    3,
			expected: Student{3, "Vivek", "hsr"},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address"}).AddRow(3, "Vivek", "hsr"))},
		},
		{
			desc:     "Case 3",
			input:    2,
			expected: Student{2, "Nithin", "Tumkur"},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address"}).AddRow(2, "Nithin", "Tumkur"))},
		},
		{
			desc:     "Case 4",
			input:    1,
			expected: Student{1, "Tejas", "Blore"},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address"}).AddRow(1, "Tejas", "Blore"))},
		},
	}

	for _, tcs := range testCasesForGetStudent {
		result, _ := dbHandler.SelectStudentById(tcs.input)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}

func TestUpdateStudent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := &SqlDB{db}

	if err != nil {
		t.Error(err)
	}

	query := "update student set name=?, address=? where id=?"

	testCasesForUpdateStudent := []struct {
		desc     string
		Student  Student
		err      error
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			Student:  Student{7, "yash", "delhi"},
			expected: errors.New("FAILED TO UPDATE STUDENT"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(7, "yash", "delhi").WillReturnError(errors.New("FAILED TO UPDATE STUDENT"))},
		},
		{
			desc:     "Case 2",
			Student:  Student{-33, "suhas", "mysore"},
			expected: errors.New("INVALID STUDENT ID"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(-33, "suhas", "mysore").WillReturnError(errors.New("INVALID STUDENT ID"))},
		},
	}

	for _, tcs := range testCasesForUpdateStudent {
		output := dbHandler.UpdateStudent(tcs.Student.id, tcs.Student.name, tcs.Student.address)
		if !reflect.DeepEqual(output, tcs.expected) {
			t.Errorf("case : %v , output: %v, expected: %v", tcs.desc, output, tcs.expected)
		}
	}
}

func TestDeleteStudent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := &SqlDB{db}
	if err != nil {
		t.Error(err)
	}
	query := "delete from student where id=?"

	testCasesForDeleteStudent := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			input:    -8,
			expected: errors.New("FAILED TO DELETE THE STUDENT"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(-8).WillReturnError(errors.New("FAILED TO DELETE THE Student"))},
		},
		{
			desc:     "Case 2",
			input:    50,
			expected: errors.New("STUDENT NOT EXISTS"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(50).WillReturnError(errors.New("STUDENT NOT EXISTS"))},
		},
		{
			desc:     "Case 3",
			input:    1,
			expected: nil,
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tcs := range testCasesForDeleteStudent {
		result := dbHandler.DeleteStudent(tcs.input)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}

func TestInsertStudent(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	dbHandler := &SqlDB{db}
	query := "insert into student(id, name, address) values(?, ?, ?)"

	testCasesForAddStudent := []struct {
		desc     string
		Student  Student
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			Student:  Student{4, "sanjay", "pune"},
			expected: errors.New("FAILED TO ADD NEW STUDENT"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(4, "sanjay", "pune").WillReturnError(errors.New("FAILED TO ADD NEW STUDENT"))},
		},
		{
			desc:     "Case 2",
			Student:  Student{-8, "sanjay", "pune"},
			expected: errors.New("STUDENT ID IS INVALID"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(-8, "sanjay", "pune").WillReturnError(errors.New("STUDENT ID IS INVALID"))},
		},
	}

	for _, tcs := range testCasesForAddStudent {
		result := dbHandler.InsertStudent(tcs.Student.id, tcs.Student.name, tcs.Student.address)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}
