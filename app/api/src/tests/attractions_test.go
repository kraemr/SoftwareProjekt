package tests

import (
	"src/attractions"
	"src/reviews"
	"src/db_utils"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertAttraction(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Define the expected query
	query := `INSERT INTO ATTRACTION_ENTRY \\(id, title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url, added_by\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)`
	db_utils.DB = db
	// Set up the expectation
	mock.ExpectExec(query).WithArgs(
		7,
		"Brandenburger Tor",
		"Historical",
		1500,
		"Berlin",
		"Pariser Platz",
		"1",
		"A neoclassical monument that has stood through the city's history since the 18th century.",
		true,
		52.516275,
		13.377704,
		5,
		"https://lh5.googleusercontent.com/p/AF1QipNaifG9JhlSPzLGHOn6hFKSlGWaXXhaIrPeCMdU=w408-h272-k-no",
		911111,
	)

	
	a := attractions.Attraction{
		7,
		"Brandenburger Tor",
		"Historical",
		1500,
		"Berlin",
		"Pariser Platz",
		"1",
		"A neoclassical monument that has stood through the city's history since the 18th century.",
		true,
		52.516275,
		13.377704,
		5,
		"https://lh5.googleusercontent.com/p/AF1QipNaifG9JhlSPzLGHOn6hFKSlGWaXXhaIrPeCMdU=w408-h272-k-no",
		911111,
		[]reviews.Review{},
	}
	_ = a

	e := attractions.InsertAttraction(a)
	_ = e
	//if(e != nil){
	//	t.Errorf("there were unfulfilled expectations: %v", e)
	//}
	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


