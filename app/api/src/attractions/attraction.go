package attractions;
import (
	_ "errors"
	"database/sql"
	"src/db_utils"
	"src/reviews"
	"fmt"
)

type Attraction struct{
	Id   			  int64	    `json:id`
	Title 			  string    `json:"title"`
	Type  			  string    `json:"type"` 
	Recommended_count int 	    `json:"recommended_count"`
	City 			  string    `json:"city"`
	Street		      string    `json:street`
	Housenumber		  string 	`json:housenumber`
	Info 			  string    `json:"info"`
	Approved		  bool	    `json:approved`		
	PosX 			  float32   `json:"posX"`
	PosY 			  float32   `json:"posY"`
	Stars			  float32   `json:stars`
	Img_url			  string    `json:img_url`
	Reviews			  []reviews.Review  `json:reviews`
}

type Filter struct{
	Filter_on string		`json:"filter_on"`
	Filter_by string		`json:"filter_by"`
	Filter_value string		`json:"filter_value"`
}


var ErrNoAttraction = fmt.Errorf("No Attractions Found")


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
	prepared_stmt,err := db.Prepare("INSERT INTO ATTRACTION_ENTRY(title,type,recommended_count,city,street,housenumber,info,PosX,PosY,stars,img_url) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if(err != nil){
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result,err := prepared_stmt.Exec(a.Title,a.Type,a.Recommended_count,a.City,a.Street,a.Housenumber,a.Info,a.PosX,a.PosY,a.Stars,a.Img_url)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

// This function Also Makes sure that the City string is made to be lowercase
func UpdateAttraction(a Attraction) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("UPDATE ATTRACTION_ENTRY SET title=?,type=?,recommended_count=?,city=?,street=?,housenumber=?,info=?,PosX=?,PosY=?,img_url=? WHERE id=?")
	if(err != nil){
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result,err := prepared_stmt.Exec(a.Title,a.Type,a.Recommended_count,a.City,a.Street,a.Housenumber,a.Info,a.PosX,a.PosY,a.Img_url,a.Id)
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
	row, err := db.Query("SELECT id,title,type,recommended_count,city,street,housenumber,info,approved,PosX,PosY,stars,img_url FROM ATTRACTION_ENTRY WHERE id = ?", id)
	
	if(err != nil){
		return Attraction{},err
	}
	defer row.Close()
	var a Attraction;

	nodata_found := true
	for row.Next() {
		row.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		nodata_found = false
	}	

	if(err != nil){
		return Attraction{},err
	}else if(nodata_found){
		return Attraction{},ErrNoAttraction
	}

	return a,nil
}

var ErrNoRecommendation = fmt.Errorf("No Recommendations Found")
func GetRecommendationForUser(id int32,city string,pref_type string) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var recommended_attractions []Attraction 
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE type = ? and city = ? ORDER BY stars LIMIT 4", pref_type,city)
	if(err != nil){
		return recommended_attractions,err
	}
	defer rows.Close()

	nodata_found := true
	var a Attraction 
	for rows.Next() {
		nodata_found = false
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		recommended_attractions = append(recommended_attractions, a)
	}	
	if(nodata_found){
		return nil,ErrNoRecommendation 
	}else if(err != nil){
		return nil,err
	}
	return recommended_attractions,nil
}

func GetAttractions() ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY")
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()
	nodata_found := true
	for rows.Next(){
		nodata_found = false
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		attractions = append(attractions, a)
	}	
	if(err != nil){
		return attractions,err
	}else if(nodata_found){
		return nil,ErrNoAttraction
	}
	return attractions,nil

}

func GetAttractionsByPos(posx float32,posy float32) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction 
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE PosX=? and PosY=?", posx,posy)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()
	nodata_found := true

	for rows.Next() {
		nodata_found = false
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}else if(nodata_found){
		return nil,ErrNoAttraction
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
	nodata_found := true


	for rows.Next() {
		nodata_found = false
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}else if(nodata_found){
		return nil,ErrNoAttraction
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
	nodata_found := true
	for rows.Next() {
		nodata_found = false
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}else if(nodata_found){
		return nil,ErrNoAttraction
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
	_ = db
	var attractions []Attraction
	title_like_str := "%" + title + "%"
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE title LIKE ? LIMIT 1000", title_like_str)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()
	nodata_found := true

	for rows.Next() {
		nodata_found = false
		a := Attraction{};
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Street,&a.Housenumber,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars,&a.Img_url)
		attractions = append(attractions, a)
	}	

	if(err != nil){
		return attractions,err
	}else if(nodata_found){
		return nil,ErrNoAttraction
	}
	return attractions,nil
}


// TODO: Implement some Kind of Filtering as generic as possible
func GetAttractionsWithFilters(filters []Filter)([]Attraction,error){
	return nil,nil
}

