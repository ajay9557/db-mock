package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	//"github.com/golang/mock/gomock"

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
		desc string
		id   int
		//query          string
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
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		/*{
			desc: "Case:2",
			id:   15,
			//query:          "Select * from Users where Id = ?",
			expectedError:  nil,
			expectedOutput: &User{},
			mock: []interface{}{
				//mock.ExpectQuery(Select * from Users where Id = ?).WithArgs(9).WillReturnError(errors.New("error,")),
				mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(15).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(0, "", 0, "", false)),
			},
		},*/
		// {
		// 	desc: "Case:3",
		// 	id:   9,
		// 	//query:          "Select from User swhere Id = ?",
		// 	expectedError:  errors.New("Error in Query"),
		// 	expectedOutput: nil,
		// 	mock: []interface{}{
		// 		mock.ExpectQuery("Select from Users where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
		// 	},
		// },
	}

	for _, tcs := range testcases {

		//fmt.Println("gfhh")
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(tcs.expectedOutput.Id, tcs.expectedOutput.Name, tcs.expectedOutput.Age, tcs.expectedOutput.Address, tcs.expectedOutput.Del))
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnError(errors.New("conn"))

		// t.Run(tcs.desc, func(t *testing.T) {

		resp, err := u.ReadByID(tcs.id)
		//fmt.Println("gfhh")
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
		// })

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
				mock.ExpectExec("INSERT INTO Users").WithArgs("Rohit", 34, "Whitefield, Bangalore", false).WillReturnResult(sqlmock.NewResult(1, 1)),
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},

		/*{
			desc: "Case:2",
			//value: User{0,"Rohit",34,"Whitefield, Bangalore",false},
			//query:          "Select from User swhere Id = ?",
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("INSERT INTO Users(Name, Age , Address ,Del ) values('Rohit',34,'Whitefield, Bangalore','0')").WithArgs().WillReturnError(errors.New("Error in Query")),
			},
		},*/
	}

	for _, tcs := range testcases {

		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(tcs.expectedOutput.Id, tcs.expectedOutput.Name, tcs.expectedOutput.Age, tcs.expectedOutput.Address, tcs.expectedOutput.Del))
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnError(errors.New("conn"))

		// t.Run(tcs.desc, func(t *testing.T) {

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

		//fmt.Printf("Expected Affected %v Got %v Expected LastInsertId %v Got %v \n", tcs.expectedAffected, resp1, tcs.expectedLastInsertId, resp2)
		// })

	}
}

func Test_Read(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc string
		//query          string
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
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		/*{

		   {
			desc: "Case:2",
			id:   9,
			//query:          "Select from User swhere Id = ?",
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select from Users where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
			},
		},*/
	}

	for _, tcs := range testcases {

		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(tcs.expectedOutput.Id, tcs.expectedOutput.Name, tcs.expectedOutput.Age, tcs.expectedOutput.Address, tcs.expectedOutput.Del))
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnError(errors.New("conn"))

		// t.Run(tcs.desc, func(t *testing.T) {

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
		// })

	}
}

func Test_Update(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc  string
		value string
		id    int
		//query          string
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
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		{
			desc:                 "Case:1",
			value:                "Jack",
			id:                   65,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("Update Users Set Name = ? where Id = ?").WithArgs("Jack", 65).WillReturnResult(sqlmock.NewResult(1, 1)),
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		/*{
			desc: "Case:3",
			id:   9,
			//query:          "Select from User swhere Id = ?",
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select from Users where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
			},
		},*/
	}

	for _, tcs := range testcases {

		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(tcs.expectedOutput.Id, tcs.expectedOutput.Name, tcs.expectedOutput.Age, tcs.expectedOutput.Address, tcs.expectedOutput.Del))
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnError(errors.New("conn"))

		// t.Run(tcs.desc, func(t *testing.T) {

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
		// })

	}
}

func Test_Delete(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc string
		//value string
		id int
		//query          string
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc: "Case:1",
			//value:                "Jack",
			id:                   6,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM Users WHERE Id = ?").WithArgs(6).WillReturnResult(sqlmock.NewResult(1, 1)),
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		{
			desc: "Case:1",
			//value:                "Jack",
			id:                   36,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM Users WHERE Id = ?").WithArgs(36).WillReturnResult(sqlmock.NewResult(1, 1)),
				//mock.ExpectQuery("Select * from Users where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(5, "Karun", 20, "HSR, Bangalore", false)),
			},
		},
		/*{
			desc: "Case:3",
			id:   9,
			//query:          "Select from User swhere Id = ?",
			expectedError:  errors.New("Error in Query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select from Users where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in Query")),
			},
		},*/
	}

	for _, tcs := range testcases {

		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address", "Del"}).AddRow(tcs.expectedOutput.Id, tcs.expectedOutput.Name, tcs.expectedOutput.Age, tcs.expectedOutput.Address, tcs.expectedOutput.Del))
		//mock.ExpectQuery(tcs.query).WithArgs(tcs.id).WillReturnError(errors.New("conn"))

		// t.Run(tcs.desc, func(t *testing.T) {

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
		// })

	}
}
