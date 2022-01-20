package sql

import (
	"database/sql"
	_"log"
	_"errors"
	_"fmt"
	"context"

	_ "github.com/go-sql-driver/mysql"
)

// repository represent the repository model
type sqlDb struct {
	db *sql.DB
}
type user struct {
	id    int
	name  string 
	email string 
	phone int 
}
// func New(db *sql.DB)*user {
// 	return &user {
// 		db: db,
// 	}
// }

// NewRepository will create a variable that represent the Repository struct
// func NewConnection() ( *sqlDb, error) {
// 	db, err := sql.Open("mysql", "root:password@localhost(3630)/test")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &sqlDb{db}, nil
// }

// Close attaches the provider and close the connection

// FindByID attaches the user repository and find data based on id
func (d *sqlDb) FindByID(id int) (*user,error) {

//	defer db.Close()

	res, err := d.db.Query("SELECT id, name, email, phone FROM users WHERE id = ?", id)

	if err != nil {
		return nil,err
	}
	defer res.Close()

	var user1 user
	for res.Next() {
		err := res.Scan(&user1.id, &user1.name, &user1.email, &user1.phone)
		if err!=nil{
			return nil, err
		}
	}

	return &user1, nil
	
}


func (d *sqlDb) Create(id int, name string ,email string, phone int) (int,error) {
	query := "insert into users(id, name, email, phone) values (?, ?)"
	resp, err := d.db.ExecContext(context.TODO(), query, id, name, email, phone)
	if err != nil {
		return 0, err
	}
	userId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(userId), nil
}

// Update attaches the user repository and update data based on id
func (DB *sqlDb) Update(id int, name string, email string, phone int) error {
	query := "update users set name = ? where id = ? "
	resp, err := DB.db.ExecContext(context.TODO(), query, name, id)
	if err != nil {
		return err
	}
	_, err = resp.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}


// Find attaches the user repository and find all data
// func Find() (error) {
// 	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/test")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	res, err := db.Query("select * from student")

// 	if err != nil {
// 		panic(err)
// 	}

// 	for res.Next() {
// 		var user1 user
// 		err := res.Scan(&user1.id, &user1.name, &user1.email, &user1.phone)

// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return nil
// }

// Create attaches the user repository and creating the data



// Delete attaches the user repository and delete data based on id


// func (DB *sqlDb)Delete(id int, name string, email string, phone int) error {
// 	query := "delete from student where id = ?"
// 	resp, err := DB.db.ExecContext(context.TODO(), query, id)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = resp.rowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
