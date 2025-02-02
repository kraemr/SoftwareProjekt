package tests

import (
	"src/db_utils"
	"src/favorites"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDeleteAttractionFavoriteById(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	id := int64(1)

	// Expecting a prepared statement for deleting a favorite by ID and the execution of the prepared statement
	prep := mock.ExpectPrepare("DELETE FROM USER_FAVORITE WHERE id=\\?")
	prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))

	if err := favorites.DeleteAttractionFavoriteById(id); err != nil {
		t.Errorf("Error was not expected while deleting attraction favorite: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAttractionFavoriteByAttractionId(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	attraction_id := int64(2)
	// Expecting a prepared statement for deleting a favorite by ID and the execution of the prepared statement
	prep := mock.ExpectPrepare("DELETE FROM USER_FAVORITE WHERE id=\\?")
	prep.ExpectExec().WithArgs(attraction_id).WillReturnResult(sqlmock.NewResult(0, 1))

	if err := favorites.DeleteAttractionFavoriteByAttractionId(attraction_id); err != nil {
		t.Errorf("Error was not expected while deleting attraction favorite by attraction id: %s", err)
	}

	// We check whether the expectation that one row should have been affected is met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestAddAttractionFavoriteById(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	userID := int64(1)
	attractionID := int64(100)
	// Expecting a prepared statement for adding a favorite by ID and the execution of the prepared statement
	prep := mock.ExpectPrepare("INSERT INTO USER_FAVORITE\\(user_id,attraction_id\\) VALUES\\(\\?,\\?\\)")
	prep.ExpectExec().WithArgs(userID, attractionID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := favorites.AddAttractionFavoriteById(userID, attractionID); err != nil {
		t.Errorf("Error was not expected while adding attraction favorite: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAttractionFavoritesByUserId(t *testing.T) {
	// Mocking the database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting the mocked database connection
	db_utils.DB = db

	userID := int64(1)
	// Mocking the rows that will be returned by the query
	rows := sqlmock.NewRows([]string{"id", "user_id", "attraction_id", "attraction_id", "title", "type", "recommended_count", "city", "info", "approved", "posX", "posY", "stars"}).
		AddRow(1, userID, 101, 101, "Attraction One", "Type A", 10, "City X", "Information", true, 12.34, 56.78, 4.5)

	// Expecting a query to retrieve the favorites by user ID and the return of the mocked rows
	mock.ExpectQuery("SELECT \\* FROM USER_FAVORITE as uf JOIN ATTRACTION_ENTRY as at ON uf.attraction_id = at.id WHERE user_id=\\?").
		WithArgs(userID).
		WillReturnRows(rows)

	favs, err := favorites.GetAttractionFavoritesByUserId(userID)
	if err != nil {
		t.Errorf("Error was not expected while retrieving favorites: %s", err)
	}
	if len(favs) == 0 {
		t.Errorf("Expected to retrieve at least one favorite")
	}

	// Test for the case where no favorites are found
	mock.ExpectQuery("SELECT \\* FROM USER_FAVORITE as uf JOIN ATTRACTION_ENTRY as at ON uf.attraction_id = at.id WHERE user_id=\\?").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err = favorites.GetAttractionFavoritesByUserId(userID)
	if err == nil {
		t.Errorf("Expected an error when no favorites are found")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
