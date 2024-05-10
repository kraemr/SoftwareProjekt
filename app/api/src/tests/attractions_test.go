package tests
import (
	"testing"
	"src/attractions"
)

//UNIT-TEST ATTRACTION
func TestInsertAttraction(t *testing.T){
	test_attraction := attractions.Attraction{}
	test_attraction.Title = "testTitle"
	test_attraction.Type = "testType"
	test_attraction.Info = "testInfo"
	test_attraction.Recommended_count = 100000
	test_attraction.PosX = 20.0
	test_attraction.PosY = 8.542
	err := attractions.InsertAttraction(test_attraction)

	if(err != nil){
		t.Fatalf(`InsertAttraction(): %v`,err)
	}

}

func TestGetAttraction(t *testing.T){

}

func TestGetAttractionByName(t *testing.T){
}

func TestGetAttractionCount(t *testing.T){
}

func TestGetAttractionsWithFilters(t *testing.T){

}
//UNIT-TEST ATTRACTION
