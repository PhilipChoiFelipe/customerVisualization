package users

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

//TODO: add tests for the various functions in user.go, as described in the assignment.
//use `go test -cover` to ensure that you are covering all or nearly all of your code paths.

func TestValidate(t *testing.T) {
	cases := []struct {
		testName     string
		testUser     NewUser
		expectsError bool
	}{
		{
			"Valid User information",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "min", LastName: "choie",
			},
			false,
		},
		{
			"Wrong Email Format",
			NewUser{
				Email: "wrongEmailFormat", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "min", LastName: "choie",
			},
			true,
		},
		{
			"Wrong Password length",
			NewUser{
				Email: "felipe@gmail.com", Password: "short", PasswordConf: "short",
				UserName: "felipe", FirstName: "min", LastName: "choie",
			},
			true,
		},
		{
			"Password and PasswordConf different",
			NewUser{
				Email: "felipe@gmail.com", Password: "differentPass", PasswordConf: "PassDifferent",
				UserName: "felipe", FirstName: "min", LastName: "choie",
			},
			true,
		},
		{
			"Username is empty",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "", FirstName: "min", LastName: "choie",
			},
			true,
		},
		{
			"Username contains space between letters",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "fe l i pe", FirstName: "min", LastName: "choie",
			},
			true,
		},
	}
	for _, c := range cases {
		err := c.testUser.Validate()
		if err != nil && !c.expectsError {
			t.Errorf("case %s: unexpected error: %v", c.testName, err)
		}
		if err == nil && c.expectsError {
			t.Errorf("case %s: expected error but did not get one", c.testName)
		}
	}
}

func TestToUser(t *testing.T) {
	newUser := NewUser{
		Email: "fElipe@GmaIl.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
		UserName: "felipe", FirstName: "min", LastName: "choie",
	}
	user, err := newUser.ToUser()
	if err != nil {
		t.Errorf("error in validating new user: %v", err)
	}
	if user.PhotoURL != gravatarBasePhotoURL+"45f17208dd064e4f04151db4488a40ad" {
		t.Errorf("photoURL generated incorrectly")
	}
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(newUser.Password)); err != nil {
		t.Errorf("password hashed incorrectly")
	}
}

func TestFullName(t *testing.T) {
	cases := []struct {
		testName      string
		testUser      NewUser
		expectingName string
	}{
		{
			"Full Name",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "Min", LastName: "Choi",
			},
			"Min Choi",
		},
		{
			"FirstName empty",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "", LastName: "Choi",
			},
			"Choi",
		},
		{
			"LastName empty",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "Min", LastName: "",
			},
			"Min",
		},
		{
			"Both Name empty",
			NewUser{
				Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
				UserName: "felipe", FirstName: "", LastName: "",
			},
			"",
		},
	}
	for _, c := range cases {
		user, err := c.testUser.ToUser()
		if err != nil {
			t.Errorf("error in validating new user: %v", err)
		}
		if user.FullName() != c.expectingName {
			t.Errorf("expected name %s, but got %s", c.expectingName, user.FullName())
		}
	}
}

func TestAuthenticate(t *testing.T) {
	newUser := NewUser{
		Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
		UserName: "felipe", FirstName: "Min", LastName: "Choi",
	}
	cases := []struct {
		testName string
		value    string
		correct  bool
	}{
		{
			"Correct Password",
			newUser.Password,
			true,
		},
		{
			"Wrong Password",
			"wrongPassword",
			false,
		},
		{
			"Emptry Password",
			"",
			false,
		},
	}
	user, err := newUser.ToUser()
	if err != nil {
		t.Errorf("error in validating new user: %v", err)
	}
	for _, c := range cases {
		err := user.Authenticate(c.value)
		if err == nil && c.correct == false {
			t.Errorf("password should be wrong but returned no error")
		}
		if err != nil && c.correct == true {
			t.Errorf("password should be equal but returned error: %v", err)
		}
	}
}

func TestApplyUpdates(t *testing.T) {
	updates := &Updates{FirstName: "ChangedFirst", LastName: "ChangedLast"}
	newUser := NewUser{
		Email: "felipe@gmail.com", Password: "1q2w3e4r!@", PasswordConf: "1q2w3e4r!@",
		UserName: "felipe", FirstName: "Min", LastName: "Choi",
	}
	user, err := newUser.ToUser()
	if err != nil {
		t.Errorf("error in validating new user: %v", err)
	}
	if user.ApplyUpdates(updates); user.FirstName != "ChangedFirst" || user.LastName != "ChangedLast" {
		t.Errorf("user names should be changed but did not change correctly")
	}
}
