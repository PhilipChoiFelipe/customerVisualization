package users

import (
	"errors"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//ErrUserNotInserted is returned when user can't be inserted
var ErrUserNotInserted = errors.New("failed to insert new User")

//ErrUserNotUpdated is returned when user can't be updated
var ErrUserNotUpdated = errors.New("failed to update user")

//ErrUserNotDeleted is returned when user can't be deleted
var ErrUserNotDeleted = errors.New("failed to delete user")

//Storage represents a databse communicating store for Users
type Storage interface {
	//GetByID returns the User with the given ID
	GetById(id int64) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *Updates) (*User, error)

	//Delete deletes the user with the given ID
	Delete(id int64) error

	//Insert signin history to UserSignIn
	InsertSignIn(id int64, clientIP string) error

	//Logs out signin history of given userId
	LogSignIns(id int64) error
}
