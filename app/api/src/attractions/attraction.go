package attractions

import (
	"database/sql"
	_ "errors"
	"fmt"
	"src/db_utils"
	"src/reviews"
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
	Added_by		  int32    `json:added_by`
	Reviews			  []reviews.Review  `json:reviews`
}

type Filter struct{
	Filter_on string		`json:"filter_on"`
	Filter_by string		`json:"filter_by"`
	Filter_value string		`json:"filter_value"`
}

var ErrNoAttraction = fmt.Errorf("No Attractions Found")
var ErrNotModerator = fmt.Errorf("Not a Moderator")

func getAttractionsFromDb(rows *sql.Rows) ([]Attraction, error) {
	var attr_list []Attraction
	no_data := true
	for rows.Next() {
		no_data = false
		a := Attraction{}
		rows.Scan(
			&a.Id,
			&a.Title,
			&a.Type,
			&a.Recommended_count,
			&a.City,
			&a.Street,
			&a.Housenumber,
			&a.Info,
			&a.Approved,
			&a.PosX,
			&a.PosY,
			&a.Stars,
			&a.Img_url,
			&a.Added_by)
		attr_list = append(attr_list, a)
	}
	if no_data {
		return nil, ErrNoAttraction
	}
	return attr_list, nil
}

func RemoveAttraction(id int64) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("DELETE FROM ATTRACTION_ENTRY WHERE id = ?")
	if err != nil {
		fmt.Println("Couldnt Remove Attraction")
		return err
	}
	result, err := prepared_stmt.Exec(id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

// This function Also Makes sure that the City string is made to be lowercase
func InsertAttraction(a Attraction) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("INSERT INTO ATTRACTION_ENTRY(title,type,recommended_count,city,street,housenumber,info,PosX,PosY,stars,img_url,added_by) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result, err := prepared_stmt.Exec(a.Title, a.Type, a.Recommended_count, a.City, a.Street, a.Housenumber, a.Info, a.PosX, a.PosY, a.Stars, a.Img_url, a.Added_by)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

// This function Also Makes sure that the City string is made to be lowercase
func UpdateAttraction(a Attraction) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("UPDATE ATTRACTION_ENTRY SET title=?,type=?,recommended_count=?,city=?,street=?,housenumber=?,info=?,PosX=?,PosY=?,img_url=? WHERE id=?")
	if err != nil {
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result, err := prepared_stmt.Exec(a.Title, a.Type, a.Recommended_count, a.City, a.Street, a.Housenumber, a.Info, a.PosX, a.PosY, a.Img_url, a.Id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func ChangeAttractionApproval(approved bool, id int64) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("UPDATE ATTRACTION_ENTRY SET ATTRACTION_ENTRY.approved=? WHERE ATTRACTION_ENTRY.id=?")
	if err != nil {
		fmt.Println("Couldnt Create Approve/Disapprove Attraction PreparedSTMT")
		return err
	}
	result, err := prepared_stmt.Exec(approved, id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func GetAttraction(id int64) (Attraction, error) {
	var db *sql.DB = db_utils.DB
	row, err := db.Query("SELECT id,title,type,recommended_count,city,street,housenumber,info,approved,PosX,PosY,stars,img_url,added_by FROM ATTRACTION_ENTRY WHERE id = ? and approved=TRUE", id)
	if(err != nil){
		return Attraction{},err
	}
	defer row.Close()
	var a Attraction

	nodata_found := true
	for row.Next() {
		row.Scan(&a.Id, &a.Title, &a.Type, &a.Recommended_count, &a.City, &a.Street, &a.Housenumber, &a.Info, &a.Approved, &a.PosX, &a.PosY, &a.Stars, &a.Img_url, &a.Added_by)
		nodata_found = false
	}

	if err != nil {
		return Attraction{}, err
	} else if nodata_found {
		return Attraction{}, ErrNoAttraction
	}

	return a, nil
}

var ErrNoRecommendation = fmt.Errorf("No Recommendations Found")

func GetRecommendationForUser(id int32, city string, pref_type string) ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	var recommended_attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE type = ? and city = ? and approved=TRUE ORDER BY stars LIMIT 4", pref_type, city)
	if err != nil {
		return recommended_attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

func GetAttractions() ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY where approved=TRUE")
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

func GetAttractionsAddedBy(user_id int32) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY Where added_by = ?",user_id)
	if(err != nil){
		return attractions,err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}



func GetAttractionsUnapprovedCity(city string) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE city = ? and approved=FALSE", city)
	if err != nil {
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

func GetAttractionsByPos(posx float32, posy float32) ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE PosX=? and PosY=? and approved=TRUE", posx, posy)
	if err != nil {
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

func GetAttractionsByCategory(category string) ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE type = ? and approved=TRUE", category)
	if err != nil {
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

// Get Attraction By City String where City is converted to lowercase always
func GetAttractionsByCity(city string) ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE city = ? and approved=TRUE", city)
	if err != nil {
		fmt.Println(err.Error())
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}

/*
Gets Attraction by Everything Similiar
Example User Types a
Finds everything starting with a
*/
func GetAttractionsByTitle(title string) ([]Attraction, error) {
	var db *sql.DB = db_utils.DB
	_ = db
	var attractions []Attraction
	title_like_str := "%" + title + "%"
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE title LIKE ? and approved=TRUE LIMIT 1000", title_like_str)
	if err != nil {
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}
