package attractions;
import (
	_ "errors"
	"database/sql"
	"fmt"
)

var(
	db *sql.DB
)

type Attraction struct{
	Title string 			`json:"title"`
	Type string  			`json:"type"` 
	Recommended_count int 	`json:"recommended_count"`
	City string				`json:"city"`
	Info string				`json:"info"`
	PosX float32			`json:"posX"`
	PosY float32			`json:"posY"`
}

type Filter struct{
	Filter_on string		`json:"filter_on"`
	Filter_by string		`json:"filter_by"`
	Filter_value string		`json:"filter_value"`
}

func InsertAttraction(a Attraction) error{
	prepared_stmt,err := db.Prepare("INSERT INTO ATTRACTION_ENTRY(title,type,recommended_count,city,info,PosX,PosY) VALUES(?,?,?,?,?,?,?)")
	if(err != nil){
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	prepared_stmt.Exec(a.Title,a.Type,a.Recommended_count,a.Info,a.PosX,a.PosY)
	if err != nil {
		return err
	}
	return nil
}


func GetAttraction(id int) Attraction{



	return Attraction{}
}

func GetAttractionByName(name string) Attraction{
	return Attraction{}
}

func GetAttractionCount() int{
	return 1
}

func GetAttractionsWithFilters(filters []Filter)([]Attraction,error){
	return nil,nil
}

