package tests

import (
	_ "src/attractions"
	_ "src/db_utils"
	_ "testing"
)

/*

func TestRemoveAttraction(t *testing.T) {
	db_utils.InitDB()

	// Insert a sample attraction to be removed.
	testAttraction := attractions.Attraction{
		Title:    "Temporary Attraction",
		Type:     "Test",
		Added_by: 1,
	}
	_ = attractions.InsertAttraction(testAttraction)

	// Test removing the attraction.
	err := attractions.RemoveAttraction(testAttraction.Id)
	if err != nil {
		t.Fatalf("RemoveAttraction() failed: %v", err)
	}

	// Test removing a non-existing attraction.
	err = attractions.RemoveAttraction(99999999)
	if err == nil {
		t.Fatal("RemoveAttraction() should fail when trying to remove non-existing attraction")
	}
}

// UNIT-TEST ATTRACTION
func TestInsertAttraction(t *testing.T) {
	db_utils.InitDB()
	test_attraction := attractions.Attraction{}
	test_attraction.Id = 1921831
	test_attraction.Title = "testTitle"
	test_attraction.Type = "testType"
	test_attraction.Info = "testInfo"
	test_attraction.Recommended_count = 100000
	test_attraction.PosX = 20.0
	test_attraction.PosY = 8.542
	test_attraction.Stars = 1.00
	err := attractions.RemoveAttraction(1921831)
	if err != nil {
		t.Fatalf(`InsertAttraction(): %v`, err)
	}
	err = attractions.InsertAttraction(test_attraction)
	if err != nil {
		t.Fatalf(`InsertAttraction(): %v`, err)
	}
}

func TestUpdateAttraction(t *testing.T) {
	db_utils.InitDB()

	// Insert a sample attraction to be updated.
	testAttraction := attractions.Attraction{
		Id:                int64(20),
		Title:             "Old Title",
		Type:              "Old Type",
		Recommended_count: 0,
		City:              "Mainz",
		Street:            "Test Street",
		Housenumber:       "1",
		Info:              "Test Info",
		Approved:          false,
		PosX:              0.0,
		PosY:              0.0,
		Stars:             0.0,
		Img_url:           "",
		Added_by:          0,
		Reviews:           nil,
	}

	_ = attractions.InsertAttraction(testAttraction)

	// Update the attraction.
	testAttraction.Title = "New Title"
	testAttraction.Type = "New Type"
	err := attractions.UpdateAttraction(testAttraction)
	if err != nil {
		t.Fatalf("UpdateAttraction() failed: %v", err)
	}

	// Verify the updates.
	updatedAttr, _ := attractions.GetAttraction(testAttraction.Id)
	if updatedAttr.Title != "New Title" || updatedAttr.Type != "New Type" {
		t.Fatal("UpdateAttraction() failed to update fields correctly")
	}

	// Test updating a non-existing attraction.
	testAttraction.Id = 99999999
	err = attractions.UpdateAttraction(testAttraction)
	if err == nil {
		t.Fatal("UpdateAttraction() should fail when trying to update non-existing attraction")
	}
}

func TestChangeAttractionApproval(t *testing.T) {
	db_utils.InitDB()

	// Insert a sample attraction.
	testAttraction := attractions.Attraction{
		Title:    "Approval Test",
		Approved: false,
	}
	_ = attractions.InsertAttraction(testAttraction)

	// Change approval status.
	err := attractions.ChangeAttractionApproval(true, testAttraction.Id)
	if err != nil {
		t.Fatalf("ChangeAttractionApproval() failed: %v", err)
	}

	// Verify the change.
	updatedAttr, _ := attractions.GetAttraction(testAttraction.Id)
	if !updatedAttr.Approved {
		t.Fatal("ChangeAttractionApproval() failed to update approval status")
	}

	// Test changing approval of a non-existing attraction.
	err = attractions.ChangeAttractionApproval(true, 99999999)
	if err == nil {
		t.Fatal("ChangeAttractionApproval() should fail when trying to update non-existing attraction")
	}
}

func TestGetAttaction(t *testing.T) {
	db_utils.InitDB()

	// Test getting attractions when none are available.
	attrs, err := attractions.GetAttractions()
	if err != nil || len(attrs) != 0 {
		t.Fatalf("GetAttractions() should return an empty slice with no error, got %v, error: %v", len(attrs), err)
	}

	// Insert some sample attractions.
	_ = attractions.InsertAttraction(attractions.Attraction{Title: "Attraction 1", Approved: true})
	_ = attractions.InsertAttraction(attractions.Attraction{Title: "Attraction 2", Approved: true})

	// Test getting attractions.
	attrs, err = attractions.GetAttractions()
	if err != nil || len(attrs) != 2 {
		t.Fatalf("GetAttractions() failed, expected 2 attractions, got %d, error: %v", len(attrs), err)
	}
}

/*
func TestGetAttractionsByCity(t *testing.T) {
    db_utils.InitDB()

    // Stellen Sie sicher, dass keine Attraktionen zurückgegeben werden, wenn die Datenbank leer ist.
    attractions, err := attractions.GetAttractionsByCity("Berlin")
    if err != nil || len(attractions) != 0 {
        t.Fatalf("GetAttractionsByCity() sollte eine leere Liste ohne Fehler zurückgeben, wenn keine Attraktionen vorhanden sind, erhalten: %d, Fehler: %v", len(attractions), err)
    }

    // Fügen Sie genehmigte Attraktionen in verschiedenen Städten hinzu.
    _ = attractions.InsertAttraction(attractions.Attraction{Title: "Brandenburger Tor", City: "Berlin", Approved: true})
    _ = attractions.InsertAttraction(attractions.Attraction{Title: "Reichstag", City: "Berlin", Approved: true})
    _ = attractions.InsertAttraction(attractions.Attraction{Title: "Dom", City: "Köln", Approved: true})

    // Testen Sie den Abruf von Attraktionen in einer Stadt mit mehreren Einträgen.
    attractions, err = attractions.GetAttractionsByCity("Berlin")
    if err != nil || len(attractions) != 2 {
        t.Fatalf("GetAttractionsByCity() sollte zwei Attraktionen für 'Berlin' zurückgeben, erhalten: %d, Fehler: %v", len(attractions), err)
    }

    // Überprüfen Sie die spezifischen zurückgegebenen Attraktionen.
    titles := map[string]bool{"Brandenburger Tor": false, "Reichstag": false}
    for _, attr := range attractions {
        if _, exists := titles[attr.Title]; exists {
            titles[attr.Title] = true
        } else {
            t.Errorf("Unexpected attraction '%s' returned", attr.Title)
        }
    }
    for title, found := range titles {
        if !found {
            t.Errorf("Expected attraction '%s' not returned", title)
        }
    }

    // Testen Sie den Abruf von Attraktionen in einer Stadt ohne genehmigte Einträge.
    attractions, err = attractions.GetAttractionsByCity("München")
    if err != nil || len(attractions) != 0 {
        t.Fatalf("GetAttractionsByCity() sollte eine leere Liste für 'München' zurückgeben, erhalten: %d, Fehler: %v", len(attractions), err)
    }

    // Testen Sie die Groß- und Kleinschreibung der Stadtnamen.
    attractions, err = attractions.GetAttractionsByCity("berlin")
    if err != nil || len(attractions) != 2 {
        t.Fatalf("GetAttractionsByCity() sollte Groß- und Kleinschreibung ignorieren und zwei Attraktionen für 'berlin' zurückgeben, erhalten: %d, Fehler: %v", len(attractions), err)
    }
} */

/*
func TestGetAttractionsByCategory(t *testing.T) {
	attr, err := attractions.GetAttractionsByCategory("AFSNAFJASNKFNSAKJFNKAJSFNKJASNFJKANSKJFNASJFNAKJS")
	if err == nil {
		t.Fatalf(`GetAttractionsByCategory(): SHOULD RETURN AN ERROR ON NON EXISTANT DATA`)
	}
	_ = attr
}

func TestGetAttractionTitle(t *testing.T) {
	attr, err := attractions.GetAttractionsByTitle("AFSNAFJASNKFNSAKJFNKAJSFNKJASNFJKANSKJFNASJFNAKJS")
	if err == nil {
		t.Fatalf(`GetAttractionsByTitle(): SHOULD RETURN AN ERROR ON NON EXISTANT DATA`)
	}
	_ = attr
}

*/
