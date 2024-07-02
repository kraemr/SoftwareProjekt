package tests

import (
	"src/db_utils"
	_ "src/attractions"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)





func TestInsertAttraction(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	db_utils.DB = db
	mock.ExpectExec("INSERT INTO ATTRACTION_ENTRY (id,title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url,added_by) VALUES (1,'Brandenburg Gate', 'Monument', 12000, 'Berlin', 'Pariser Platz', '1', 'An 18th-century neoclassical monument in Berlin, one of the most well-known landmarks of Germany.', TRUE, 52.5163, 13.3777, 4.7, 'https://example.com/brandenburg_gate.jpg',911111),(2,'Neuschwanstein Castle', 'Castle', 11000, 'Schwangau', 'Neuschwansteinstraße', '20', 'A 19th-century Romanesque Revival palace on a rugged hill above the village of Hohenschwangau near Füssen in southwest Bavaria.', TRUE, 47.5576, 10.7498, 4.8, 'https://example.com/neuschwanstein_castle.jpg',911111),(3,'Cologne Cathedral', 'Cathedral', 13000, 'Cologne', 'Domkloster', '4', 'A renowned monument of German Catholicism and Gothic architecture and is a World Heritage Site.', TRUE, 50.9413, 6.9583, 4.9, 'https://example.com/cologne_cathedral.jpg',911111),(4,'Heidelberg Castle', 'Castle', 9500, 'Heidelberg', 'Schlosshof', '1', 'A famous ruin in Germany and landmark of Heidelberg.', TRUE, 49.4106, 8.7153, 4.6, 'https://example.com/heidelberg_castle.jpg',911111);").
	WillReturnResult(sqlmock.NewResult(1, 1))
	


}


