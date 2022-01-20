package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//db, mock, error := sqlmock.New()
	dbHandler := &SqlDB{db}
	query := "select * from userdemo1 where id=?"

	testCasesForGetUser := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			input:    4,
			expected: User{4, "Tipesh", "Three", 87},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "age"}).AddRow(4, "Tipesh", "Three", 87))},
		},
		{
			desc:     "Case 2",
			input:    3,
			expected: User{3, "Dipesh", "Two", 43},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "age"}).AddRow(3, "Dipesh", "Two", 43))},
		},
		{
			desc:     "Case 3",
			input:    8,
			expected: User{8, "Love", "Shove", 35},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(8).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "age"}).AddRow(8, "Love", "Shove", 35))},
		},
		{
			desc:     "Case 4",
			input:    10,
			expected: User{0, "", "", 0},
			mock:     []interface{}{mock.ExpectQuery(query).WithArgs(10).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "age"}).AddRow(0, "", "", 0))},
		},
	}

	for _, tcs := range testCasesForGetUser {
		result, _ := dbHandler.GetUser(tcs.input)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}

//test case two

func Test_DeleteUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//db, mock, error := sqlmock.New()
	dbHandler := &SqlDB{db}
	query := "delete from userdemo1 where id=?"

	testCasesForDeleteUser := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			input:    1,
			expected: nil,
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
		{
			desc:     "Case 2",
			input:    101,
			expected: errors.New("FAILED TO DELETE THE USER"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(1).WillReturnError(errors.New("FAILED TO DELETE THE USER"))},
		},
		{
			desc:     "Case 3",
			input:    88,
			expected: errors.New("ID NOT FOUND"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(88).WillReturnError(errors.New("ID NOT FOUND"))},
		},
	}

	for _, tcs := range testCasesForDeleteUser {
		result := dbHandler.DeleteUser(tcs.input)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}

//TEST CASE THREE
func Test_AddUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//db, mock, error := sqlmock.New()
	dbHandler := &SqlDB{db}
	query := "insert into userdemo1(id, firstname, lastname, age) values(?, ?, ?, ?)"

	testCasesForAddUser := []struct {
		desc     string
		user     User
		expected interface{}
		mock     []interface{}
	}{
		{
			desc:     "Case 1",
			user:     User{11, "Ram", "Sharma", 54},
			expected: errors.New("FAILED TO ADD USER"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(11, "Ram", "Sharma", 22).WillReturnError(errors.New("FAILED TO ADD USER"))},
		},
		{
			desc:     "Case 2",
			user:     User{109, "Ram", "Sharma", 54},
			expected: errors.New("INVALID ID"),
			mock:     []interface{}{mock.ExpectExec(query).WithArgs(11, "Ram", "Sharma", 22).WillReturnError(errors.New("INVALID ID"))},
		},
	}

	for _, tcs := range testCasesForAddUser {
		result := dbHandler.AddUser(tcs.user.id, tcs.user.firstName, tcs.user.lastName, tcs.user.age)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("output: %v, expected: %v", result, tcs.expected)
		}
	}
}

//test case four
func Test_UpdateUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//db, mock, error := sqlmock.New()
	dbHandler := &SqlDB{db}
	query := "update userdemo1 set firstname=?, lastname=?, age=? where id=?"

	testCasesForUpdateUser := []struct {
		desc     string
		user     User
		err      error
		expected interface{}
		mock     []interface{}
	}{
		// {
		// 	desc:     "Case 1",
		// 	user:     User{11, "Ram", "Sharma", 54},
		// 	expected: errors.New("FAILED TO UPDATE USER"),
		// 	mock:     []interface{}{mock.ExpectExec(query).WithArgs(11, "Ram", "Sharma", 22).WillReturnError(errors.New("FAILED TO UPDATE USER"))},
		// },
		{
			desc:     "Case 2",
			user:     User{3, "Ram", "Sharma", 54},
			expected: nil,
			mock:     []interface{}{mock.ExpectExec(query).WithArgs("Ram", "Sharma", 54, 3).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tcs := range testCasesForUpdateUser {
		result := dbHandler.UpdateUser(tcs.user.id, tcs.user.firstName, tcs.user.lastName, tcs.user.age)
		if !reflect.DeepEqual(result, tcs.expected) {
			t.Errorf("case : %v , output: %v, expected: %v", tcs.desc, result, tcs.expected)
		}
	}
}
