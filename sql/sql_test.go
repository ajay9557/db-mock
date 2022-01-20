package sql

import (
	"database/sql"
	"log"
	"testing"
	_"errors"

	"github.com/DATA-DOG/go-sqlmock"
	//	"github.com/google/uuid"
	_"github.com/stretchr/testify/assert"
	reflect "reflect"
)

// var u = []user{
// 	{
// 		id:    "1",
// 		name:  "Prasath",
// 		email: "prasath@gmail.com",
// 		phone: "9884697050",
// 	},
// }

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
//	s := New(db)
    s := sqlDb{db:db}
	query := "SELECT id, name, email, phone FROM users WHERE id = ?"

	tcs := []struct {
		testcase int
		id int
		expectedErr error
		expectedOut *user
	}{
		{
			testcase: 1,
			id: 1,
			expectedErr: nil,
			expectedOut: &user{id : 1, name:"prasath", email: "prasath@gmail.com", phone: 9884697050,},

		},
	}

	for _,tc := range tcs {
		mock.ExpectQuery(query).WithArgs(tc.id).WillReturnRows(sqlmock.NewRows([]string{"id","name","email","phone"}).AddRow(1,"prasath","prasath@gmail.com",9884697050))
		resp, err := s.FindByID(tc.id)
		if !reflect.DeepEqual(resp, tc.expectedOut) {
			t.Errorf("TestCase[%v] Expected :\t %v\nGot: \t%v\n", tc.testcase, tc.expectedOut,resp)
		}
		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("TestCase[%v] Expected : \t %v\nGot: \t%v\n", tc.testcase,tc.expectedOut,err)
		}
	}
}





func Test_Create(t *testing.T) {
	db, mock := NewMock()
//	s := New(db)
    s := sqlDb{db:db}
	query := "insert into users(id, name) values (?, ?)"

	tests := []struct {
		testCase int
		userName string
		userId   int
		usermail string
		userphone int
		resp     int
		err      error
		mock     []interface{}
	}{
		{
			testCase: 1,
			userName: "prasath",
			userId:   1,
			usermail: "prasath@gmail.com",
			userphone:  988467050,
			resp:     1,
			err:      nil,
			mock: []interface{} { mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},	
}

	for _, tc := range tests {
		resp, err := s.Create(tc.userId, tc.userName, tc.usermail,tc.userphone)
		if !reflect.DeepEqual(resp, tc.resp) {
			t.Errorf("TestCase[%v] Expected: \t%v\n Got: \t%v\n", tc.testCase, tc.resp, resp)
		}
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.err, err)
		}
	} 
}


func Test_Update(t *testing.T) {
	db, mock := NewMock()
//	s := New(db)
    s := sqlDb{db:db}
	query := "update users set name = ? where id = ? "

	tests := []struct {
		testCase int
		userName string
		userId   int
		resp     int
		err      error
		mock     []interface{}
	}{
		{
			testCase: 1,
			userName: "test_new",
			userId:   1,
			resp:     1,
			err:      nil,
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs("test_new", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	for _, tc := range tests {
		err := s.Update(tc.userId, tc.userName, tc.usermail, tc.userphone)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.err, err)
		}
	}

}




// func TestFind(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &sqlDb{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := "SELECT id, name, email, phone FROM users"

// 	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
// 	AddRow(u[0].id, u[0].name, u[0].email, u[0].phone)

// 	mock.ExpectQuery(query).WillReturnRows(rows)

// 	users, err := repo.Find()
// 	assert.NotEmpty(t, users)
// 	assert.NoError(t, err)
// 	assert.Len(t, users, 1)
// }

