package tests

import (
	"reflect"
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

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	// Setup the mock expectations for a successful update
	mock.ExpectExec("UPDATE USER SET email=\\?, password=\\?, city=\\?, username=\\? WHERE id=\\?").
		WithArgs("new@example.com", "newpassword", "NewCity", "newUsername", 1).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulating 1 row affected

	// Test updating an existing user
	user := users.User{UserId: 1, Email: "new@example.com", Password: "newpassword", City: "NewCity", Username: "newUsername"}
	if err := users.UpdateUser(user); err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	// Setup the mock expectations for an update attempt on a non-existing user
	mock.ExpectExec("UPDATE USER SET email=\\?, password=\\?, city=\\?, username=\\? WHERE id=\\?").
		WithArgs("new@example.com", "newpassword", "NewCity", "newUsername", 999).
		WillReturnResult(sqlmock.NewResult(0, 0)) // Simulating 0 rows affected

	// Test updating a non-existing user
	user = users.User{UserId: 999, Email: "new@example.com", Password: "newpassword", City: "NewCity", Username: "newUsername"}
	if err := users.UpdateUser(user); err == nil {
		t.Errorf("error was expected while updating non-existing user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	// Expected to delete an existing user
	mock.ExpectExec("DELETE FROM USER WHERE id=\\?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := users.DeleteUser(1); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}

	// Expected not to find the user
	mock.ExpectExec("DELETE FROM USER WHERE id=\\?").
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(0, 0))

	if err := users.DeleteUser(999); err == nil {
		t.Errorf("error was expected when no user exists to delete")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	rows := sqlmock.NewRows([]string{"id", "email", "password", "city", "username"}).
		AddRow(1, "test@example.com", "password123", "TestCity", "testUser")

	mock.ExpectQuery("SELECT id, email, password, city, username FROM USER WHERE id=\\? LIMIT 1").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := users.GetUserByID(1)
	if err != nil {
		t.Errorf("error was not expected while fetching user: %s", err)
	}

	expected := users.User{UserId: 1, Email: "test@example.com", Password: "password123", City: "TestCity", Username: "testUser"}
	if user != expected {
		t.Errorf("expected %v, got %v", expected, user)
	}

	// Testing the no user found scenario
	mock.ExpectQuery("SELECT id, email, password, city, username FROM USER WHERE id=\\? LIMIT 1").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = users.GetUserByID(999)
	if err == nil {
		t.Errorf("expected ErrNoUser error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserCityById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	rows := sqlmock.NewRows([]string{"city"}).
		AddRow("TestCity")

	mock.ExpectQuery("SELECT city from USER WHERE id=\\? LIMIT 1").
		WithArgs(1).
		WillReturnRows(rows)

	city, err := users.GetUserCityById(1)
	if err != nil {
		t.Errorf("error was not expected while fetching user city: %s", err)
	}
	if city != "TestCity" {
		t.Errorf("expected city to be 'TestCity', got '%s'", city)
	}

	// Testing the no city found scenario
	mock.ExpectQuery("SELECT city from USER WHERE id=\\? LIMIT 1").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows(nil))

	city, err = users.GetUserCityById(999)
	if err == nil || city != "" {
		t.Errorf("expected error and empty city for non-existing user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsersByCityAndBanned(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	// Vorbereiten der Mock-Daten
	rows := sqlmock.NewRows([]string{"UserId", "Email", "Password", "City", "Username", "Activated"}).
		AddRow(1, "user@example.com", "password", "CityX", "username1", "FALSE")

	// Erwartete SQL-Abfrage
	mock.ExpectQuery("SELECT UserId, Email, Password, City, Username, Activated FROM USER WHERE city=\\? AND activated='FALSE'").
		WithArgs("CityX").
		WillReturnRows(rows)

	// Testfunktion ausführen
	userList, err := users.GetUsersByCityAndBanned("CityX")
	if err != nil {
		t.Errorf("Error was not expected while fetching data: %s", err)
	}

	// Erwartete Nutzerdaten
	expected := []users.User{
		{UserId: 1, Email: "user@example.com", Password: "password", City: "CityX", Username: "username1", Activated: "FALSE"},
	}

	if len(userList) != 1 || !reflect.DeepEqual(userList[0], expected[0]) {
		t.Errorf("Expected %v, got %v", expected, userList)
	}

	// Überprüfen, ob alle Erwartungen erfüllt wurden
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
