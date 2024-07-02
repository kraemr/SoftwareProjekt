package tests

import (
	"src/db_utils"
	"src/moderator"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetModeratorById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	var exampleId int64 = 1
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

	email := "example@example.com"
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

	// Test when the moderator does not exist
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

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

	// Test when no moderators are found in the city
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db_utils.DB = db

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
