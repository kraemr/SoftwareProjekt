package tests

import (
	"fmt"
	"src/attractions"
	"src/db_utils"
	"src/reviews"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertAttraction(t *testing.T) {
	// Mocking the database
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Set expected SQL Prepared Statement
	prep := mock.ExpectPrepare("INSERT INTO ATTRACTION_ENTRY\\(title,type,recommended_count,city,street,housenumber,info,PosX,PosY,stars,img_url,added_by\\) VALUES\\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)")

	prep.ExpectExec().WithArgs(
		"Test Attraction",
		"Museum",
		100,
		"berlin",
		"Main St",
		"123",
		"Very interesting place",
		sqlmock.AnyArg(), // Für posX, posY und stars verwenden wir AnyArg oder einen ähnlichen Matcher
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		"http://example.com/image.png", 42,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	// Set up the expectation

	attraction := attractions.Attraction{
		Title:             "Test Attraction",
		Type:              "Museum",
		Recommended_count: 100,
		City:              "berlin",
		Street:            "Main St",
		Housenumber:       "123",
		Info:              "Very interesting place",
		PosX:              52.5, // Beachte hier die Verwendung von Fließkommazahlen
		PosY:              13.4,
		Stars:             5.0,
		Img_url:           "http://example.com/image.png",
		Added_by:          42,
		Reviews:           []reviews.Review{},
	}

	if err := attractions.InsertAttraction(attraction); err != nil {
		t.Errorf("error was not expected while inserting attraction: %s", err)
	}

	// Überprüfe alle gesetzten Erwartungen
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveAttraction(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()
	prep := mock.ExpectPrepare("DELETE FROM ATTRACTION_ENTRY WHERE id = \\?")
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := attractions.RemoveAttraction(1); err != nil {
		t.Errorf("error was not expected while inserting attraction: %s", err)
	}

	// Überprüfe alle gesetzten Erwartungen
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateAttraction(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()
	prep := mock.ExpectPrepare("UPDATE ATTRACTION_ENTRY SET title=\\?,type=\\?,recommended_count=\\?,city=\\?,street=\\?,housenumber=\\?,info=\\?,PosX=\\?,PosY=\\?,img_url=\\? WHERE id=\\?")

	prep.ExpectExec().WithArgs(
		"Test Attraction",
		"Museum",
		100,
		"berlin",
		"Main St",
		"123",
		"Very interesting place",
		sqlmock.AnyArg(), // Für posX, posY und stars verwenden wir AnyArg oder einen ähnlichen Matcher
		sqlmock.AnyArg(),
		"http://example.com/image.png",
		1,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	// Set up the expectation

	attraction := attractions.Attraction{
		Id:                1,
		Title:             "Test Attraction",
		Type:              "Museum",
		Recommended_count: 100,
		City:              "berlin",
		Street:            "Main St",
		Housenumber:       "123",
		Info:              "Very interesting place",
		PosX:              52.5, // Beachte hier die Verwendung von Fließkommazahlen
		PosY:              13.4,
		Stars:             5.0,
		Img_url:           "http://example.com/image.png",
		Added_by:          42,
		Reviews:           []reviews.Review{},
	}

	if err := attractions.UpdateAttraction(attraction); err != nil {
		t.Errorf("error was not expected while inserting attraction: %s", err)
	}

	// Überprüfe alle gesetzten Erwartungen
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangeAttractionApproval(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()
	prep := mock.ExpectPrepare("UPDATE ATTRACTION_ENTRY SET ATTRACTION_ENTRY.approved=\\? WHERE ATTRACTION_ENTRY.id=\\?")
	prep.ExpectExec().WithArgs(
		true,
		1,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	// Set up the expectation

	if err := attractions.ChangeAttractionApproval(true, 1); err != nil {
		t.Errorf("error was not expected while inserting attraction: %s", err)
	}

	// Überprüfe alle gesetzten Erwartungen
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAttraction(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("SELECT id,title,type,recommended_count,city,street,housenumber,info,approved,PosX,PosY,stars,img_url,added_by FROM ATTRACTION_ENTRY WHERE id = \\? and approved=TRUE").
		WithArgs(17).
		WillReturnRows(rowsExpected)

	attr, e := attractions.GetAttraction(17)
	if e != nil {
		t.Errorf("error")
	}
	_ = attr
	_ = e
	if err := mock.ExpectationsWereMet(); err != nil {
		fmt.Println(err)
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetAttractions(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY where approved=TRUE$").
		WillReturnRows(rowsExpected)

	attr, e := attractions.GetAttractions()
	_ = attr
	_ = e
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAttractionsAddedBy(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY Where added_by = \\?$").
		WithArgs(911111).
		WillReturnRows(rowsExpected)

	a, e := attractions.GetAttractionsAddedBy(911111)
	_ = a
	if e != nil {

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetAttractionsUnapprovedCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY WHERE city = \\? and approved=FALSE$").
		WithArgs("Berlin").
		WillReturnRows(rowsExpected)

	a, e := attractions.GetAttractionsUnapprovedCity("Berlin")
	_ = a
	if e != nil {

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAttractionsByCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY WHERE type = \\? and approved=TRUE$").
		WithArgs("Museum").
		WillReturnRows(rowsExpected)

	a, e := attractions.GetAttractionsByCategory("Museum")
	_ = a
	if e != nil {

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Get Attraction By City String where City is converted to lowercase always
func TestGetAttractionsByCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) from ATTRACTION_ENTRY WHERE city = \\? and approved=TRUE$").
		WithArgs("Berlin").
		WillReturnRows(rowsExpected)

	a, e := attractions.GetAttractionsByCity("Berlin")
	_ = a
	if e != nil {

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAttractionsByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	db_utils.DB = db
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
		AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111)

	mock.ExpectQuery("^SELECT (.+) from ATTRACTION_ENTRY WHERE title LIKE \\? and approved=TRUE LIMIT 1000$").
		WithArgs("%ergamon%"). // getAttractionsByTitle adds % around the string
		WillReturnRows(rowsExpected)

	a, e := attractions.GetAttractionsByTitle("ergamon")
	_ = a
	if e != nil {

	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
