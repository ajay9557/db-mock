package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock1() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error %s in opening the Database connection", err)

	}
	return db, mock
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("Error %s in opening the Database connection", err)

	}
	return db, mock
}

func New(db *sql.DB) *DbUser {

	return &DbUser{db: db}
}

func Test_ReadByID(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc           string
		id             int
		expectedError  error
		expectedOutput *User
		mock           []interface{}
	}{
		{
			desc:           "Case:1",
			id:             5,
			expectedError:  nil,
			expectedOutput: &User{Id: 5, Name: "Karun", Age: 20, Address: "HSR, Bangalore", Del: false},
			mock: []interface{}{
				mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		{
			desc:           "Case:2",
			id:             15,
			expectedError:  nil,
			expectedOutput: &User{},
			mock: []interface{}{

				mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(15).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(0, "", 0, "", false)),
			},
		},
		{
			desc:           "Case:3",
			id:             9,
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp, err := u.ReadByID(tcs.id)

		if !reflect.DeepEqual(resp, tcs.expectedOutput) {
			t.Errorf("Expected %v Got %v\n", tcs.expectedOutput, resp)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

		fmt.Printf("Expected %v Got %v\n", tcs.expectedOutput, resp)

	}
}

func Test_Create(t *testing.T) {

	db, mock := NewMock1()
	u := New(db)

	testcases := []struct {
		desc                 string
		value                User
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			value:                User{0, "Rohit", 34, "Whitefield, Bangalore", false},
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO Users").WithArgs("Rohit", 34, "Whitefield, Bangalore", false).WillReturnResult(sqlmock.NewResult(1, 1))},
		},

		{
			desc:                 "Case:2",
			expectedError:        errors.New("Error in Query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO Users").WithArgs("Itachi", 34, "Whitefield, Bangalore", false).WithArgs().WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Create(tcs.value)
		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}

func Test_Read(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc           string
		expectedError  error
		expectedOutput *User
		mock           []interface{}
	}{
		{
			desc:           "Case:1",
			expectedError:  nil,
			expectedOutput: &User{Id: 5, Name: "Karun", Age: 20, Address: "HSR, Bangalore", Del: false},
			mock: []interface{}{
				mock.ExpectQuery("Select * from Users").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		{

			desc:           "Case:2",
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from Users").WithArgs().WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp, err := u.Read()
		if !reflect.DeepEqual(resp, tcs.expectedOutput) {
			t.Errorf("Expected %v Got %v\n", tcs.expectedOutput, resp)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

		fmt.Printf("Expected %v Got %v\n", tcs.expectedOutput, resp)

	}
}

func Test_Update(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc                 string
		value                string
		id                   int
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			value:                "Jack",
			id:                   5,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("Update Users Set Name = ? where Id = ?").WithArgs("Jack", 5).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:                 "Case:2",
			value:                "Jack",
			id:                   65,
			expectedError:        nil,
			expectedLastInsertId: 0,
			expectedAffected:     0,
			mock: []interface{}{
				mock.ExpectExec("Update Users Set Name = ? where Id = ?").WithArgs("Jack", 65).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
		{
			desc:                 "Case:3",
			id:                   9,
			value:                "Jack",
			expectedError:        errors.New("Error in Query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("Update Users Set Name = ? where Id = ?").WithArgs("Jack", 9).WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Update(tcs.value, tcs.id)

		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}

func Test_Delete(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc                 string
		id                   int
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			id:                   6,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM Users WHERE Id = ?").WithArgs(6).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:                 "Case:2",
			id:                   36,
			expectedError:        nil,
			expectedLastInsertId: 0,
			expectedAffected:     0,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM Users WHERE Id = ?").WithArgs(36).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
		{
			desc:                 "Case:3",
			id:                   9,
			expectedError:        errors.New("Error in Query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM Users WHERE Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Delete(tcs.id)

		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}
