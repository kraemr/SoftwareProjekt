package tests

import (
	"src/db_utils"
	"src/users"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	mock.ExpectExec("INSERT INTO USER \\(email, password, city, username\\)").
		WithArgs("test@example.com", "password123", "TestCity", "testUser").
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := users.User{Email: "test@example.com", Password: "password123", City: "TestCity", Username: "testUser"}
	if err := users.CreateUser(user); err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	rows := sqlmock.NewRows([]string{"id", "email", "password", "city", "username"}).
		AddRow(1, "test@example.com", "password123", "TestCity", "testUser")

	mock.ExpectQuery("SELECT id from USER WHERE email=\\? LIMIT 1").
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id, email, password, city, username from USER WHERE id=\\? LIMIT 1").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := users.GetUserByEmail("test@example.com")
	if err != nil {
		t.Errorf("error was not expected while fetching user: %s", err)
	}

	expected := users.User{UserId: 1, Email: "test@example.com", Password: "password123", City: "TestCity", Username: "testUser"}
	if user != expected {
		t.Errorf("expected %v, got %v", expected, user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
