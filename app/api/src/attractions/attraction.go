package attractions;
import (
	_ "errors"
	"database/sql"
	"src/db_utils"
	"fmt"
)

type Review struct{
	Text string `json:text`
	Username string `json:username`
	Userid int64 `json:userid` 
	Stars float32 `json:stars`
}

type Attraction struct{
	Id   			  int64	   `json:id`
	Title 			  string   `json:"title"`
	Type  			  string   `json:"type"` 
	Recommended_count int 	   `json:"recommended_count"`
	City 			  string   `json:"city"`
	Info 			  string   `json:"info"`
	Approved		  bool	   `json:approved`		
	PosX 			  float32  `json:"posX"`
	PosY 			  float32  `json:"posY"`
	Stars			  float32  `json:stars`
	Reviews			  []Review `json:reviews`
}

type Filter struct{
	Filter_on string		`json:"filter_on"`
	Filter_by string		`json:"filter_by"`
	Filter_value string		`json:"filter_value"`
}



func RemoveAttraction(id int64) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("DELETE FROM ATTRACTION_ENTRY WHERE id = ?")
	if(err != nil){
		fmt.Println("Couldnt Remove Attraction")
		return err
	}
	result,err := prepared_stmt.Exec(id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

// This function Also Makes sure that the City string is made to be lowercase
func InsertAttraction(a Attraction) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("INSERT INTO ATTRACTION_ENTRY(title,type,recommended_count,city,info,PosX,PosY,stars) VALUES(?,?,?,?,?,?,?,?)")
	if(err != nil){
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result,err := prepared_stmt.Exec(a.Title,a.Type,a.Recommended_count,a.City,a.Info,a.PosX,a.PosY,a.Stars)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

// This function Also Makes sure that the City string is made to be lowercase
func UpdateAttraction(a Attraction) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("UPDATE ATTRACTION_ENTRY SET title=?,type=?,recommended_count=?,city=?,info=?,PosX=?,PosY=? WHERE id=?")
	if(err != nil){
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result,err := prepared_stmt.Exec(a.Title,a.Type,a.Recommended_count,a.City,a.Info,a.PosX,a.PosY,a.Id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func ChangeAttractionApproval(approved bool,id int64) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("UPDATE ATTRACTION_ENTRY SET ATTRACTION_ENTRY.approved=? WHERE ATTRACTION_ENTRY.id=?")
	if(err != nil){
		fmt.Println("Couldnt Create Approve/Disapprove Attraction PreparedSTMT")
		return err
	}
	result,err := prepared_stmt.Exec(approved,id)
	_ = result
	if err != nil {
		return err
	}
	return nil

}

func GetAttraction(id int) (Attraction,error){
	var db *sql.DB = db_utils.DB
	row, err := db.Query("SELECT id,title,type,recommended_count,city,info,approved,PosX,PosY,stars FROM ATTRACTION_ENTRY WHERE id = ?", id)
	
	if(err != nil){
		return Attraction{},err
	}
	defer row.Close()
	var a Attraction;

	for row.Next() {
		row.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
	}	

	if(err != nil){
		return Attraction{},err
	}
	return a,nil
}

func GetAttractionsByPos(posx float32,posy float32) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction 
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE PosX=? and PosY=?", posx,posy)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()

	for rows.Next() {
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}
	return attractions,nil
}


func GetAttractionsByCategory(category string) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction 
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE type = ?", category)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()

	for rows.Next() {
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}
	return attractions,nil
}

// Get Attraction By City String where City is converted to lowercase always
func GetAttractionsByCity(city string) ( []Attraction,error) {
	var db *sql.DB = db_utils.DB
	var attractions []Attraction 
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE city = ?", city)
	if(err != nil){
		fmt.Println(err.Error())
		return attractions,err
	}
	defer rows.Close()

	for rows.Next() {
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}
	return attractions,nil
}

/* 
Gets Attraction by Everything Similiar
Example User Types a
Finds everything starting with a
*/
func GetAttractionsByTitle(title string) ( []Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	title_like_str := "%" + title + "%"
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE title LIKE ? LIMIT 1000", title_like_str)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()

	for rows.Next() {
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}
	return attractions,nil
}


// TODO: Implement some Kind of Filtering as generic as possible
func GetAttractionsWithFilters(filters []Filter)([]Attraction,error){
	return nil,nil
}

