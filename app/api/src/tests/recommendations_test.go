package tests;
import (
	"src/db_utils"
	"src/recommendations"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetRecommendations(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    db_utils.DB = db
	
	rowsExpected := 
	sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
	AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111).
    AddRow(22, "Topography of Terror", "Museum", 450, "Berlin", "Niederkirchnerstraße", "8", "An indoor and outdoor museum documenting the terror of the Nazi regime.", true, 52.5063, 13.3849, 4.5, "https://lh5.googleusercontent.com/p/AF1QipPZ6UT85ELKmYOcfU59m6L_S_MyMexdY3U_vxx4=w408-h306-k-no", 911111).
    AddRow(24, "Bode Museum", "Museum", 350, "Berlin", "Bodestraße", "1-3", "Part of the Museum Island complex, featuring collections of sculptures, coins, and Byzantine art.", true, 52.5225, 13.3953, 4, "https://lh5.googleusercontent.com/p/AF1QipPE7vtnT3z8Ks0DpV_xlnQ7SXETF_Q8LdCa8TpW=w408-h306-k-no", 911111).
    AddRow(26, "Neue Nationalgalerie", "Museum", 250, "Berlin", "Potsdamer Straße", "50", "A museum for modern art, showcasing works from the early 20th century.", true, 52.5071, 13.3654, 4.5, "https://lh5.googleusercontent.com/p/AF1QipN9_q2kyXQp_GnPlCi66tf4zlvxmfQ3JrAigsuz=w408-h408-k-no", 911111)
	
	_ = rowsExpected
	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY WHERE type = \\? and city = \\? ORDER BY stars LIMIT 4$").
	WithArgs("Museum","Berlin").
	WillReturnRows(rowsExpected)
	recomms,err1 := recommendations.GetRecommendationForUser(0,"Berlin","Museum")
	_ = recomms
	_ = err1	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("^SELECT (.+) FROM ATTRACTION_ENTRY WHERE type = \\? and city = \\? ORDER BY stars LIMIT 4$").
	WithArgs("Museum","Mainz").WillReturnRows(sqlmock.NewRows([]string{}))
	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	
}