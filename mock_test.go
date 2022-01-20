package main

import (
	
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)



func Test_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	New(db)

	mock.ExpectExec("CREATE TABLE album if not exists(id INT AUTO_INCREMENT NOT NULL,title VARCHAR(128) NOT NULL,artist VARCHAR(255) NOT NULL,price DECIMAL(5,2) NOT NULL,PRIMARY KEY (`id`))").WillReturnResult(sqlmock.NewResult(1, 1))

	if err = create(); err != nil {
		t.Errorf("ellehe:%s", err)
	}
}
func Test_Addalbum(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	New(db)

	tests := []struct {
		desc   string
		id     int
		Title  string
		Artist string
		Price  float64
		resp   int64
		err    error
		mock   []interface{}
	}{
		{
			desc:   "case-1",
			Title:  "sirivennela",
			Artist: "Anurag",
			Price:  100.09,
			resp:   1,
			err:    nil,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO album(title,artist,price) VALUES (?,?,?)").WithArgs("sirivennela", "Anurag", 100.09).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}

	
	for _, tc := range tests {
		resp, err := addAlbum(Album{Title: tc.Title, Artist: tc.Artist, Price: tc.Price})
		
		if resp != tc.resp && err != tc.err {
			t.Errorf("error was not expected while updating stats: %v", err)
		}

	}

}

func Test_update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	New(db)

	tests := []struct {
		desc string
		id   int

		resp error

		mock []interface{}
	}{
		{
			desc: "case-1",
			id:   1,
			resp: nil,

			mock: []interface{}{
				mock.ExpectExec("Update Album set Price=79.99 where id=?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	
	for _, tc := range tests {
		err := updateAlbum(tc.id)
		
		if err != tc.resp {
			t.Errorf("%v", tc.resp)
		}
	}

}

func Test_delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	New(db)

	tests := []struct {
		desc string
		id   int

		resp error

		mock []interface{}
	}{
		{
			desc: "case-1",
			id:   3,
			resp: nil,

			mock: []interface{}{
				mock.ExpectExec("Delete From Album where id=?").WithArgs(3).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	
	for _, tc := range tests {
		err := delete(tc.id)
		
		if err != tc.resp {
			t.Errorf("%v", tc.resp)
		}
	}

}

func Test_ID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	New(db)
	rows := sqlmock.NewRows([]string{"ID", "Title", "Artist", "Price"}).
		
		AddRow(1, "Sarah Vaughan", "Sarah vaughan", 34.98)
	mock.ExpectQuery("SELECT * FROM album WHERE id=?").WithArgs(1).WillReturnRows(rows)

	if _, err = albumByID(1); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

}
