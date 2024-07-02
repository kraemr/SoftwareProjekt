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
/*
	rows := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
    AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111).
    AddRow(18, "Berlin Zoological Garden", "Zoo", 650, "Berlin", "Hardenbergplatz", "8", "The oldest and most famous zoo in Germany, home to a wide variety of species.", true, 52.5075, 13.3372, 4.5, "https://lh3.googleusercontent.com/gps-proxy/ALd4DhHfMFwGMyiMylWs_B2axBgIB3tQkuw1wvPvC7rZASmTKj9teIMcxKiPB7Bfex1Ua64Y0lXx8XbEW1JN5YAmdhlDJa2NN8wEQpR0xb5UztltIoXGiDnOb6PB9SmTvFjnjhUN3qbR88-c0KqIRgbNAyzRsfpjvaJ40KGCTd4M2wq-Uj5PLyT1K8Lk=w408-h272-k-no", 911111).
    AddRow(19, "Gendarmenmarkt", "Square", 600, "Berlin", "Gendarmenmarkt", "", "A picturesque square featuring the Konzerthaus, the French Cathedral, and the German Cathedral.", true, 52.5139, 13.3924, 4.5, "https://lh5.googleusercontent.com/p/AF1QipM1Xo0cruKQ8VIIsYj7CJcgFtzh7ST0j9N85RDC=w408-h724-k-no", 911111).
    AddRow(20, "Kurfürstendamm", "Street", 550, "Berlin", "Kurfürstendamm", "", "A famous avenue known for its shops, houses, and hotels, often considered the Champs-Élysées of Berlin.", true, 52.5026, 13.3301, 4, "https://lh5.googleusercontent.com/p/AF1QipPZl-Hy7CsXOKSN1VrKdhsGxC_sW1xnGm_nQutK=w408-h544-k-no", 911111).
    AddRow(21, "Victory Column", "Monument", 500, "Berlin", "Großer Stern", "1", "A monument commemorating the Prussian victory in the Danish-Prussian War, offering a panoramic view of Berlin.", true, 52.5145, 13.3501, 4.5, "https://lh5.googleusercontent.com/p/AF1QipMnPN3c-e81mpPLTwrXxpAiITzmfK64k4GiCp6_=w408-h544-k-no", 911111).
    AddRow(22, "Topography of Terror", "Museum", 450, "Berlin", "Niederkirchnerstraße", "8", "An indoor and outdoor museum documenting the terror of the Nazi regime.", true, 52.5063, 13.3849, 4.5, "https://lh5.googleusercontent.com/p/AF1QipPZ6UT85ELKmYOcfU59m6L_S_MyMexdY3U_vxx4=w408-h306-k-no", 911111).
    AddRow(23, "Berlin TV Tower", "Observation", 400, "Berlin", "Panoramastraße", "1A", "The tallest structure in Germany, offering an observation deck with a view of Berlin.", true, 52.5208, 13.4094, 4.5, "https://lh5.googleusercontent.com/p/AF1QipMASq3OP3DrDMYRhZteXS2_Qfd6m_q8rRrqHiPH=w408-h725-k-no", 911111).
    AddRow(24, "Bode Museum", "Museum", 350, "Berlin", "Bodestraße", "1-3", "Part of the Museum Island complex, featuring collections of sculptures, coins, and Byzantine art.", true, 52.5225, 13.3953, 4, "https://lh5.googleusercontent.com/p/AF1QipPE7vtnT3z8Ks0DpV_xlnQ7SXETF_Q8LdCa8TpW=w408-h306-k-no", 911111).
    AddRow(25, "Hackesche Höfe", "Courtyard", 300, "Berlin", "Rosenthaler Straße", "40-41", "A complex of interlinked courtyards in the Spandau district, known for its vibrant cultural scene.", true, 52.5252, 13.4018, 4, "https://lh5.googleusercontent.com/p/AF1QipMV-0MTPqcHHPqEfQtutmvYTPJdj2Cl9mVd4A-H=w408-h306-k-no", 911111).
    AddRow(26, "Neue Nationalgalerie", "Museum", 250, "Berlin", "Potsdamer Straße", "50", "A museum for modern art, showcasing works from the early 20th century.", true, 52.5071, 13.3654, 4.5, "https://lh5.googleusercontent.com/p/AF1QipN9_q2kyXQp_GnPlCi66tf4zlvxmfQ3JrAigsuz=w408-h408-k-no", 911111)
*/

	rowsExpected := sqlmock.NewRows([]string{"id", "title", "type", "recommended_count", "city", "street", "housenumber", "info", "approved", "PosX", "PosY", "stars", "img_url", "added_by"}).
	AddRow(17, "Pergamon Museum", "Museum", 700, "Berlin", "Bodestraße", "1-3", "One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.", true, 52.5214, 13.3965, 5, "https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no", 911111).
    AddRow(22, "Topography of Terror", "Museum", 450, "Berlin", "Niederkirchnerstraße", "8", "An indoor and outdoor museum documenting the terror of the Nazi regime.", true, 52.5063, 13.3849, 4.5, "https://lh5.googleusercontent.com/p/AF1QipPZ6UT85ELKmYOcfU59m6L_S_MyMexdY3U_vxx4=w408-h306-k-no", 911111).
    AddRow(24, "Bode Museum", "Museum", 350, "Berlin", "Bodestraße", "1-3", "Part of the Museum Island complex, featuring collections of sculptures, coins, and Byzantine art.", true, 52.5225, 13.3953, 4, "https://lh5.googleusercontent.com/p/AF1QipPE7vtnT3z8Ks0DpV_xlnQ7SXETF_Q8LdCa8TpW=w408-h306-k-no", 911111).
    AddRow(26, "Neue Nationalgalerie", "Museum", 250, "Berlin", "Potsdamer Straße", "50", "A museum for modern art, showcasing works from the early 20th century.", true, 52.5071, 13.3654, 4.5, "https://lh5.googleusercontent.com/p/AF1QipN9_q2kyXQp_GnPlCi66tf4zlvxmfQ3JrAigsuz=w408-h408-k-no", 911111)


	mock.ExpectQuery("SELECT * FROM ATTRACTION_ENTRY WHERE type = \\? and city = \\? ORDER BY stars LIMIT 4").
	WithArgs("Museum","Berlin").
	WillReturnRows(rowsExpected);
	
	recomms,err1 := recommendations.GetRecommendationForUser(0,"Berlin","Museum")
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	
	_ = recomms
	_ = err1

	
}