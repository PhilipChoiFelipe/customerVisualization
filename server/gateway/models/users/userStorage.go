package users

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserStorage struct {
	sqlsess *sql.DB
}

//User Storage implementation

//NewSqlStorage returns new sql connection instance
func NewSqlStorage(sqlsess *sql.DB) *UserStorage {
	if sqlsess == nil {
		panic("nil sql session")
	}
	return &UserStorage{sqlsess}
}

var (
	ID        int64
	Email     string
	UserName  string
	PassHash  []byte
	FirstName string
	LastName  string
	StoreName string
)

//GetById finds id of user in DB and returns the user
func (us *UserStorage) GetById(id int64) (*User, error) {

	insq := "select id, email, username, passhash, first_name, last_name, store_name from users where id = ?"
	err := us.sqlsess.QueryRow(insq, id).Scan(&ID, &Email, &UserName, &PassHash, &FirstName, &LastName, &StoreName)
	if err != nil {
		return nil, err
	}
	return &User{ID, Email, UserName, PassHash, FirstName, LastName, StoreName}, nil
}

//GetByEmail find email of user in DB and returns the user
func (us *UserStorage) GetByEmail(email string) (*User, error) {
	insq := "select id, email, username, passhash, first_name, last_name, store_name from users where email = ?"
	err := us.sqlsess.QueryRow(insq, email).Scan(&ID, &Email, &UserName, &PassHash, &FirstName, &LastName, &StoreName)
	if err != nil {
		return nil, err
	}
	return &User{ID, Email, UserName, PassHash, FirstName, LastName, StoreName}, nil
}

//GetByUserName finds user by username and returns user
func (us *UserStorage) GetByUserName(username string) (*User, error) {
	insq := "select id, email, username, passhash, first_name, last_name, store_name from users where userName = ?"
	rows, err := us.sqlsess.Query(insq, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &Email, &UserName, &PassHash, &FirstName, &LastName, &StoreName)
	}
	return &User{ID, Email, UserName, PassHash, FirstName, LastName, StoreName}, nil
}

//Insert inserts user and returns inserted user
func (us *UserStorage) Insert(user *User) (*User, error) {
	query := "INSERT INTO users (email, username, passhash, first_name, last_name, store_name) VALUES (?,?,?,?,?,?)"
	res, err := us.sqlsess.Exec(query, user.Email, user.UserName, user.PassHash, user.FirstName, user.LastName, user.StoreName)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

//Update updates user with given user's id and update information. Returns updated user
func (us *UserStorage) Update(id int64, updates *Updates) (*User, error) {
	query := "update users set first_name = ?, last_name = ? where id = ?"
	_, err := us.sqlsess.Exec(query, updates.FirstName, updates.LastName, id)
	if err != nil {
		return nil, err
	}
	user, err := us.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//Delete deletes user of given user's id
func (us *UserStorage) Delete(id int64) error {
	insq := "delete from users where id = ?"
	_, err := us.sqlsess.Exec(insq, id)
	if err != nil {
		return err
	}
	return nil
}

//InsertSignIn inserts user of given id and ip into usersignin table
func (us *UserStorage) InsertSignIn(userID int64, clientIP string) error {
	query := "INSERT INTO UserSignIn (UserID, SignInTime, ClientIP) VALUES (?,?,?)"
	_, err := us.sqlsess.Exec(query, userID, time.Now(), clientIP)
	if err != nil {
		return err
	}
	return nil
}

//LogSignIns logs user of given id's login history
func (us *UserStorage) LogSignIns(userID int64) error {
	var (
		UserID     int64
		SignInTime time.Time
		ClientIP   string
	)
	insq := "select UserID, SignInTime, ClientIP from UserSignIn where UserID = ?"
	rows, err := us.sqlsess.Query(insq, userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&UserID, &SignInTime, &ClientIP)
		log.Println(fmt.Sprintf("Client IP: %s, User ID: %d, Signed in time: %s", ClientIP, UserID, SignInTime.String()))
	}
	return nil
}
