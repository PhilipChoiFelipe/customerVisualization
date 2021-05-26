package users

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDatabase(t *testing.T) {
	db, _, err := sqlmock.New()
	sqlStorage := NewSqlStorage(db)
	if err != nil {
		t.Errorf("error opening database: %v\n", err)
	}

	defer sqlStorage.sqlsess.Close()

	if err := sqlStorage.sqlsess.Ping(); err != nil {
		t.Errorf("error pinging database: %v\n", err)
	}
}

func TestGetById(t *testing.T) {
	cases := []struct {
		name         string
		expectedUser *User
		idToGet      int64
		expectError  bool
	}{
		{
			"User Found",
			&User{
				1,
				"test@test.com",
				[]byte("passhash123"),
				"username",
				"firstname",
				"lastname",
				"photourl",
			},
			1,
			false,
		},
		{
			"User Not Found",
			&User{},
			2,
			true,
		},
		{
			"Really Long ID",
			&User{
				1234567890,
				"test@test.com",
				[]byte("passhash123"),
				"username",
				"firstname",
				"lastname",
				"photourl",
			},
			1,
			false,
		},
	}

	query := "select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL from User where ID = ?"

	for _, c := range cases {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		row := mock.NewRows([]string{
			"ID",
			"Email",
			"PassHash",
			"UserName",
			"FirstName",
			"LastName",
			"PhotoURL"},
		).AddRow(
			c.expectedUser.ID,
			c.expectedUser.Email,
			c.expectedUser.PassHash,
			c.expectedUser.UserName,
			c.expectedUser.FirstName,
			c.expectedUser.LastName,
			c.expectedUser.PhotoURL,
		)
		sqlStorage := NewSqlStorage(db)
		if c.expectError == true {
			mock.ExpectQuery(query).WithArgs(c.idToGet).WillReturnError(ErrUserNotFound)
			user, err := sqlStorage.GetById(c.idToGet)
			if user != nil || err == nil {
				t.Errorf("Expected error [%v] but got [%v] instead", ErrUserNotFound, err)
			}
		} else {
			mock.ExpectQuery(query).WithArgs(c.idToGet).WillReturnRows(row)
			user, err := sqlStorage.GetById(c.idToGet)
			if err != nil {
				t.Errorf("Unexpected error on successful test [%s]: %v", c.name, err)
			}
			if !reflect.DeepEqual(user, c.expectedUser) {
				t.Errorf("Error, invalid match in test [%s]", c.name)
			}
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}

func TestGetByEmail(t *testing.T) {
	cases := []struct {
		name         string
		expectedUser *User
		emailToGet   string
		expectError  bool
	}{
		{
			"User Found",
			&User{
				ID:        1,
				Email:     "test@test.com",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "firstname",
				LastName:  "lastname",
				PhotoURL:  "photourl",
			},
			"test@test.com",
			false,
		},
		{
			"User Not Found: wrong email",
			&User{
				2,
				"wrong@wrong.come",
				[]byte("passhash123"),
				"username",
				"firstname",
				"lastname",
				"photourl",
			},
			"test@test.com",
			true,
		},
		{
			"User Not Found",
			&User{},
			"test@test.com",
			true,
		},
	}

	query := "select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL from User where Email = ?"

	for _, c := range cases {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		row := mock.NewRows([]string{
			"ID",
			"Email",
			"PassHash",
			"UserName",
			"FirstName",
			"LastName",
			"PhotoURL"},
		).AddRow(
			c.expectedUser.ID,
			c.expectedUser.Email,
			c.expectedUser.PassHash,
			c.expectedUser.UserName,
			c.expectedUser.FirstName,
			c.expectedUser.LastName,
			c.expectedUser.PhotoURL,
		)
		sqlStorage := NewSqlStorage(db)
		if c.expectError == true {
			mock.ExpectQuery(query).WithArgs(c.emailToGet).WillReturnError(ErrUserNotFound)
			user, err := sqlStorage.GetByEmail(c.emailToGet)
			if user != nil || err == nil {
				t.Errorf("Expected error [%v] but got [%v] instead", ErrUserNotFound, err)
			}
		} else {
			mock.ExpectQuery(query).WithArgs(c.emailToGet).WillReturnRows(row)
			user, err := sqlStorage.GetByEmail(c.emailToGet)
			if err != nil {
				t.Errorf("Unexpected error on successful test [%s]: %v", c.name, err)
			}
			if !reflect.DeepEqual(user, c.expectedUser) {
				t.Errorf("Error, invalid match in test [%s]", c.name)
			}
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}

func TestGetByUserName(t *testing.T) {
	cases := []struct {
		name          string
		expectedUser  *User
		usernameToGet string
		expectError   bool
	}{
		{
			"User Found",
			&User{
				ID:        1,
				Email:     "test@test.com",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "firstname",
				LastName:  "lastname",
				PhotoURL:  "photourl",
			},
			"username",
			false,
		},
		{
			"User Not Found: wrong username",
			&User{
				2,
				"wrong@wrong.come",
				[]byte("passhash123"),
				"wrongname",
				"firstname",
				"lastname",
				"photourl",
			},
			"wrongwrongname",
			true,
		},
		{
			"User Not Found",
			&User{},
			"NoUSER",
			true,
		},
	}

	query := "select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL from User where UserName = ?"

	for _, c := range cases {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		row := mock.NewRows([]string{
			"ID",
			"Email",
			"PassHash",
			"UserName",
			"FirstName",
			"LastName",
			"PhotoURL"},
		).AddRow(
			c.expectedUser.ID,
			c.expectedUser.Email,
			c.expectedUser.PassHash,
			c.expectedUser.UserName,
			c.expectedUser.FirstName,
			c.expectedUser.LastName,
			c.expectedUser.PhotoURL,
		)
		sqlStorage := NewSqlStorage(db)
		if c.expectError == true {
			mock.ExpectQuery(query).WithArgs(c.usernameToGet).WillReturnError(ErrUserNotFound)
			user, err := sqlStorage.GetByUserName(c.usernameToGet)
			if user != nil || err == nil {
				t.Errorf("Expected error [%v] but got [%v] instead", ErrUserNotFound, err)
			}
		} else {
			mock.ExpectQuery(query).WithArgs(c.usernameToGet).WillReturnRows(row)
			user, err := sqlStorage.GetByUserName(c.usernameToGet)
			if err != nil {
				t.Errorf("Unexpected error on successful test [%s]: %v", c.name, err)
			}
			if !reflect.DeepEqual(user, c.expectedUser) {
				t.Errorf("Error, invalid match in test [%s]", c.name)
			}
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		name        string
		insertUser  *User
		expectError bool
	}{
		{
			"User Inserted Correctly",
			&User{
				ID:        0,
				Email:     "test@test.com",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "firstname",
				LastName:  "lastname",
				PhotoURL:  "photourl",
			},
			false,
		},
		{
			"User Inserted InCorrectly: User Email is too long",
			&User{
				ID:        0,
				Email:     "test@test.comasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfcomasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfcomasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "firstname",
				LastName:  "lastname",
				PhotoURL:  "photourl",
			},
			true,
		},
	}

	query := "INSERT INTO User (Email, PassHash, UserName, FirstName, LastName, PhotoURL) VALUES (?,?,?,?,?,?)"
	for _, c := range cases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		sqlStorage := NewSqlStorage(db)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		if c.expectError == true {
			mock.ExpectExec(query).WithArgs(c.insertUser.Email, c.insertUser.PassHash, c.insertUser.UserName, c.insertUser.FirstName, c.insertUser.LastName, c.insertUser.PhotoURL).WillReturnError(ErrUserNotInserted)
			if _, err := sqlStorage.Insert(c.insertUser); err == nil {
				t.Errorf("case %s: error was expected while inserting user", c.name)
			}
		} else {
			mock.ExpectExec(query).WithArgs(c.insertUser.Email, c.insertUser.PassHash, c.insertUser.UserName, c.insertUser.FirstName, c.insertUser.LastName, c.insertUser.PhotoURL).WillReturnResult(sqlmock.NewResult(0, 1))
			if user, err := sqlStorage.Insert(c.insertUser); !reflect.DeepEqual(user, c.insertUser) {
				t.Errorf("case %s: error was not expected while inserting user: %s", c.name, err)
			}
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}

func TestUpdate(t *testing.T) {
	cases := []struct {
		name           string
		updatedUser    *User
		updatingValues *Updates
		updatingId     int64
		expectError    bool
	}{
		{
			"User updated Correctly",
			&User{
				ID:        0,
				Email:     "test@test.com",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "changedFirst",
				LastName:  "changedLast",
				PhotoURL:  "photourl",
			},
			&Updates{
				FirstName: "changedFirst",
				LastName:  "changedLast",
			},
			0,
			false,
		},
		{
			"User updated InCorrectly: names are different",
			&User{
				ID:        0,
				Email:     "test@test.com",
				PassHash:  []byte("passhash123"),
				UserName:  "username",
				FirstName: "wrongUpdate",
				LastName:  "wrongUpdate",
				PhotoURL:  "photourl",
			},
			&Updates{
				FirstName: "changedFirst",
				LastName:  "changedLast",
			},
			0,
			true,
		},
	}

	query := "update User set FirstName = ?, LastName = ? where ID = ?"
	secondQuery := "select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL from User where ID = ?"
	for _, c := range cases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		sqlStorage := NewSqlStorage(db)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		row := mock.NewRows([]string{
			"ID",
			"Email",
			"PassHash",
			"UserName",
			"FirstName",
			"LastName",
			"PhotoURL"},
		).AddRow(
			c.updatedUser.ID,
			c.updatedUser.Email,
			c.updatedUser.PassHash,
			c.updatedUser.UserName,
			c.updatedUser.FirstName,
			c.updatedUser.LastName,
			c.updatedUser.PhotoURL,
		)

		if c.expectError == true {
			mock.ExpectExec(query).WithArgs(c.updatingValues.FirstName, c.updatingValues.LastName, c.updatingId).WillReturnError(ErrUserNotUpdated)
			if _, err := sqlStorage.Update(c.updatingId, c.updatingValues); err == nil {
				t.Errorf("case %s: error was expected while updating user", c.name)
			}
		} else {
			mock.ExpectExec(query).WithArgs(c.updatingValues.FirstName, c.updatingValues.LastName, c.updatingId).WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectQuery(secondQuery).WithArgs(c.updatingId).WillReturnRows(row)
			if user, err := sqlStorage.Update(c.updatingId, c.updatingValues); !reflect.DeepEqual(user, c.updatedUser) {
				t.Errorf("case %s: error was not expected while updating user: %s", c.name, err)
			}
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}

func TestDelete(t *testing.T) {
	query := "delete User where ID = ?"
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlStorage := NewSqlStorage(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(query).WithArgs(123).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := sqlStorage.Delete(123); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
