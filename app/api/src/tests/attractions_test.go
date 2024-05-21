package tests
import (
	"testing"
	"src/attractions"
	"src/db_utils"
)

//UNIT-TEST ATTRACTION
func TestInsertAttraction(t *testing.T){
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
	if(err != nil){
		t.Fatalf(`InsertAttraction(): %v`,err)
	}
	err = attractions.InsertAttraction(test_attraction)
	if(err != nil){
		t.Fatalf(`InsertAttraction(): %v`,err)
	}
}

func TestGetNullAttraction(t *testing.T){
	attr,err := attractions.GetAttraction(93123121)
	if(err == nil){
		t.Fatalf(`GetAttraction(): SHOULD RETURN AN ERROR ON NON EXISTANT DATA`)
	}
	_ = attr
}

func TestGetNullAttractionCategory(t *testing.T){
	attr,err := attractions.GetAttractionsByCategory("AFSNAFJASNKFNSAKJFNKAJSFNKJASNFJKANSKJFNASJFNAKJS")
	if(err == nil){
		t.Fatalf(`GetAttractionsByCategory(): SHOULD RETURN AN ERROR ON NON EXISTANT DATA`)
	}
	_ = attr
} 

func TestGetNullAttractionTitle(t *testing.T){
	attr,err := attractions.GetAttractionsByTitle("AFSNAFJASNKFNSAKJFNKAJSFNKJASNFJKANSKJFNASJFNAKJS")
	if(err == nil){
		t.Fatalf(`GetAttractionsByTitle(): SHOULD RETURN AN ERROR ON NON EXISTANT DATA`)
	}
	_ = attr
}
