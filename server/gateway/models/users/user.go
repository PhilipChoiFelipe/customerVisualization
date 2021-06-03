package users

import (
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//User (business owners) represents a user account in the database
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	UserName  string `json:"userName"`
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	StoreName string `json:"storeName"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	StoreName    string `json:"storeName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserSignIn struct {
	UserID     string `json:"userID"`
	SignInTime string `json:"signInTime"`
	ClientIP   string `json:"clientIP"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	_, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return err
	}
	if len(nu.Password) < 6 {
		return fmt.Errorf("password length must be larger than 6")
	}

	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("password and confirmed password are different")
	}

	if len(nu.UserName) < 1 {
		return fmt.Errorf("user name cannot be empty")
	}

	if strings.Contains(nu.UserName, " ") {
		return fmt.Errorf("user name cannot contain space between letters")
	}

	if len(nu.StoreName) < 1 {
		return fmt.Errorf("store name cannot be empty")
	}
	return nil
}

//ToUser converts the NewUser to a User
func (nu *NewUser) ToUser() (*User, error) {
	err := nu.Validate()
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        0,
		Email:     nu.Email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		StoreName: nu.StoreName,
	}
	user.SetPassword(nu.Password)
	return user, nil
}

//FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put between the names. If both are missing,
//this returns an empty string
func (u *User) FullName() string {
	fullName := ""
	if len(u.FirstName) < 1 || len(u.LastName) < 1 {
		fullName = u.FirstName + u.LastName
	} else {
		fullName = u.FirstName + " " + u.LastName
	}
	return fullName
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	u.PassHash = passHash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	if updates.FirstName == "" || updates.LastName == "" {
		return fmt.Errorf("empty updating value")
	}
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	return nil
}
