package tests

import (
	"src/db_utils"
	"src/moderator"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetModeratorById(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	var exampleId int64 = 1
	// Expecting a query for a moderator by ID and returning a row with the moderator data
	rows := sqlmock.NewRows([]string{"id", "email", "password", "city", "username"}).
		AddRow(1, "moderator@example.com", "password123", "SampleCity", "modname")

	mock.ExpectQuery("Select \\* from CITY_MODERATOR where id = \\?").
		WithArgs(exampleId).
		WillReturnRows(rows)

	mod, err := moderator.GetModeratorById(exampleId)
	if err != nil {
		t.Errorf("Error was not expected while retrieving moderator: %s", err)
	}
	if mod.Id != 1 || mod.Email != "moderator@example.com" || mod.City != "SampleCity" || mod.Username != "modname" {
		t.Errorf("Unexpected moderator data returned: %+v", mod)
	}

	var noResultId int64 = 99
	// Testcase 2: Expecting a query for a moderator by ID and returning no rows (no moderator found)
	mock.ExpectQuery("Select \\* from CITY_MODERATOR where id = \\?").
		WithArgs(noResultId).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = moderator.GetModeratorById(noResultId)
	if err != moderator.ErrNoModerator {
		t.Errorf("Expected ErrNoModerator for no result, got: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetModeratorByEmail(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	email := "example@example.com"
	// Expecting a query for a moderator by email and returning a row with the moderator data
	rows := sqlmock.NewRows([]string{"id", "email", "password", "city", "username"}).
		AddRow(1, email, "password123", "SampleCity", "modname")

	mock.ExpectQuery("Select \\* from CITY_MODERATOR where email = \\?").
		WithArgs(email).
		WillReturnRows(rows)

	mod, err := moderator.GetModeratorByEmail(email)
	if err != nil {
		t.Errorf("Error was not expected while retrieving moderator: %s", err)
	}
	if mod.Email != email {
		t.Errorf("Unexpected moderator data returned: %+v", mod)
	}

	// Subtest: Test when no moderator is found for the email
	mock.ExpectQuery("Select \\* from CITY_MODERATOR where email = \\?").
		WithArgs("noemail@example.com").
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = moderator.GetModeratorByEmail("noemail@example.com")
	if err != moderator.ErrNoModerator {
		t.Errorf("Expected ErrNoModerator for no result, got: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetModeratorsCity(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	// Expecting a query for moderators by city and returning rows with the moderator data
	city := "SampleCity"
	rows := sqlmock.NewRows([]string{"id", "email", "password", "city", "username"}).
		AddRow(1, "moderator@example.com", "password123", city, "modname").
		AddRow(2, "moderator2@example.com", "password456", city, "modname2")

	mock.ExpectQuery("Select \\* from CITY_MODERATOR where city = \\?").
		WithArgs(city).
		WillReturnRows(rows)

	moderators, err := moderator.GetModeratorsCity(city)
	if err != nil {
		t.Errorf("Error was not expected while retrieving moderators: %s", err)
	}
	if len(moderators) != 2 {
		t.Errorf("Unexpected number of moderators returned: %d", len(moderators))
	}

	// Subtest: Test when no moderator is found for the city
	mock.ExpectQuery("Select \\* from CITY_MODERATOR where city = \\?").
		WithArgs("UnknownCity").
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = moderator.GetModeratorsCity("UnknownCity")
	if err != moderator.ErrNoModerator {
		t.Errorf("Expected ErrNoModerator for no result, got: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestUpdateModerator(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	// Expecting a prepared statement for updating a moderator and the execution of the prepared statement
	mod := moderator.Moderator{Id: 1, Email: "moderator@example.com", City: "SampleCity", Username: "modname"}
	prep := mock.ExpectPrepare("UPDATE CITY_MODERATOR SET id=\\?,email=\\?,city=\\?,username=\\? WHERE id=\\?")
	prep.ExpectExec().WithArgs(mod.Id, mod.Email, mod.City, mod.Username, mod.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	if err := moderator.UpdateModerator(mod); err != nil {
		t.Errorf("Error was not expected while updating moderator: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestInsertModerator(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	// Expecting a prepared statement for inserting a moderator and the execution of the prepared statement
	mod := moderator.Moderator{Id: 2, Email: "newmoderator@example.com", City: "NewCity", Username: "newmodname"}
	prep := mock.ExpectPrepare("INSERT INTO CITY_MODERATOR\\(id,email,city,username\\) VALUES\\(\\?,\\?,\\?,\\?\\)")
	prep.ExpectExec().WithArgs(mod.Id, mod.Email, mod.City, mod.Username).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := moderator.InsertModerator(mod); err != nil {
		t.Errorf("Error was not expected while inserting moderator: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDisableUser(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	// Expecting update query for disabling a user and the execution of the query
	email := "user@example.com"
	mock.ExpectExec("UPDATE USER SET active=false WHERE email = \\?").
		WithArgs(email).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := moderator.DisableUser(email); err != nil {
		t.Errorf("Error was not expected while disabling user: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
