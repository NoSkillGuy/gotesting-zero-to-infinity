package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID   int
    Name string
    Age  int
}

type UsersDB struct {
    Db *sql.DB
}

func InitializeDB() (*UsersDB, error) {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

   
    return &UsersDB{Db: db}, nil
}

func (u *UsersDB) CreateUsersTable() error {
    _, err := u.Db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
    return err
}

func (u *UsersDB) InsertUser(name string, age int) error {
    _, err := u.Db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
    return err
}

func (u *UsersDB) GetUserByID(id int) (*User, error) {
    row := u.Db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id)
    user := new(User)
    err := row.Scan(&user.ID, &user.Name, &user.Age)
    return user, err
}