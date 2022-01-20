package userDB

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTable(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	testCases := []struct {
		tableName   string
		tableSchema string
		mockQuery   *sqlmock.ExpectedExec
		expected    error
	}{
		{
			tableName:   "user",
			tableSchema: "id INT NOT NULL AUTO_INCREMENT,name VARCHAR(255) NOT NULL,age INT NOT NULL,deleted INT NOT NULL DEFAULT 0,PRIMARY KEY (id)",
			mockQuery: mock.ExpectExec("CREATE TABLE IF NOT EXISTS user ( id INT NOT NULL AUTO_INCREMENT,name VARCHAR(255) NOT NULL,age INT NOT NULL,deleted INT NOT NULL DEFAULT 0,PRIMARY KEY (id) )").
				WillReturnResult(sqlmock.NewResult(1, 1)), expected: nil},
		{
			tableName:   "stats",
			tableSchema: "id INT NOT NULL AUTO_INCREMEN,",
			mockQuery: mock.ExpectExec("CREATE TABLE IF NOT EXISTS user()").
				WillReturnError(errors.New("error creating the table")), expected: errors.New("error creating the table")},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := CreateTable(db, testCase.tableName, testCase.tableSchema)
			if err != nil && err.Error() != testCase.expected.Error() {
				t.Errorf("expected error: %v, got: %v", testCase.expected, err)
			}
		})
	}
}

func TestGetValuesById(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "age", "deletedFlag"}).
		AddRow(1, "John", 20, 0).AddRow(2, "Jane", 21, 0)

	testCases := []struct {
		id          int
		user        User
		mockQuery   *sqlmock.ExpectedQuery
		expectError error
	}{
		{id: 1, user: User{1, "John", 20, 0}, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(1).WillReturnRows(rows), expectError: nil},
		{id: 3, user: User{0, "", 0, 0}, mockQuery: nil, expectError: sql.ErrNoRows},
		{id: 2, user: User{2, "Jane", 21, 0}, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(2).WillReturnRows(rows), expectError: nil},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			user, err := GetValuesById(db, tableName, testCase.id)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error: %v, got: %v", testCase.expectError, err)
			}
			if !reflect.DeepEqual(user, testCase.user) {
				t.Errorf("expected user: %v, got: %v", User{1, "John", 20, 0}, user)
			}
		})
	}
}

func TestInsertValues(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer db.Close()

	testCases := []struct {
		user        User
		mockQuery   *sqlmock.ExpectedExec
		expectedErr error
	}{
		{user: User{1, "John", 20, 0}, mockQuery: mock.ExpectExec("INSERT INTO user(id, name, age) VALUES (?, ?, ?)").WithArgs(1, "John", 20).WillReturnResult(sqlmock.NewResult(1, 1)), expectedErr: nil},
		{user: User{1, "Jane", 21, 0}, mockQuery: nil, expectedErr: errors.New("error inserting values")},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := InsertValues(db, tableName, testCase.user.id, testCase.user.age, testCase.user.name)
			if err != nil && err.Error() != testCase.expectedErr.Error() {
				t.Errorf("expected error: %v, got: %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer db.Close()

	testCases := []struct {
		id       int
		mockExec *sqlmock.ExpectedExec
		expected error
	}{
		{id: 1, mockExec: mock.ExpectExec("UPDATE user SET deleted = 1 WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)), expected: nil},
		{id: 10, mockExec: mock.ExpectExec("UPDATE user SET deleted = 1 WHERE id = ?").WithArgs(1).WillReturnError(errors.New("error deleting record")), expected: errors.New("error deleting record")},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := DeleteRecord(db, tableName, testCase.id)
			if errors.Is(err, testCase.expected) {
				t.Errorf("expected error: %v, got: %v", nil, err)
			}
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tableName := "user"

	defer db.Close()

	testCases := []struct {
		id       int
		column   string
		value    interface{}
		mockExec *sqlmock.ExpectedExec
		expected error
	}{
		{
			id:       1,
			column:   "name",
			value:    "John",
			mockExec: mock.ExpectExec("UPDATE user SET name = 'John' WHERE id = 1;").WillReturnResult(sqlmock.NewResult(1, 1)),
			expected: nil},
		{
			id:       1,
			column:   "age",
			value:    20,
			mockExec: mock.ExpectExec("UPDATE user SET age = 20 WHERE id = 1;").WillReturnResult(sqlmock.NewResult(1, 1)),
			expected: nil},
		{
			id:       10,
			column:   "name",
			value:    "Jane",
			mockExec: mock.ExpectExec("UPDATE user SET name = 'Jane' WHERE id = 10;").WillReturnError(errors.New("error updating the records")),
			expected: errors.New("error updating the records")},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := UpdateRecord(db, tableName, testCase.id, testCase.column, testCase.value)
			if err != nil && err.Error() != testCase.expected.Error() {
				t.Errorf("expected error: %v, got: %v", nil, err)
			}
		})
	}
}
