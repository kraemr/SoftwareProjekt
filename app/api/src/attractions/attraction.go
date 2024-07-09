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

/*
Receives Database rows
Iterates over them and returns a list of Attractions
Also returns an error if no Attractions were found
*/
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
		fmt.Println(a)
		attr_list = append(attr_list, a)
	}
	if no_data {
		return nil, ErrNoAttraction
	}
	return attr_list, nil
}

/*
This removes an Attraction by its id from the Database
*/
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

/*
This function expects that the Attraction struct has every value except Id already
filled with values.
If that is the case a new Attraction is inserted into the Database with the given data.
Newly added Attractions will be unapproved
*/
func InsertAttraction(a Attraction) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt, err := db.Prepare("INSERT INTO ATTRACTION_ENTRY(title,type,recommended_count,city,street,housenumber,info,PosX,PosY,stars,img_url,added_by,approved) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,FALSE)")
	if err != nil {
		fmt.Println("Couldnt Insert Attraction")
		return err
	}
	result, err := prepared_stmt.Exec(a.Title, a.Type, a.Recommended_count, a.City, a.Street, a.Housenumber, a.Info, a.PosX, a.PosY, a.Stars, a.Img_url, a.Added_by)
	_ = result
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
This function expects that title,type,favorit_count,city,street,housenumber,info and so on are set in
the attraction struct, It Returns an error when the Database Exec fails.
An error indicates that values were not set properly
*/
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


/*
This function sets the attraction approval  identified by the id.
returns an error if sql fails or Attraction doesnt exist
*/
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

/*
Gets a single attraction by its Id and returns it
if there is an error it indicates that this attraction doesnt exist
*/
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


/*
Gets Recommendations For User by city and type
Ordered asending by stars and Limit to 4
*/
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

/*
Return all approved Attractions
*/
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

/*
Returns all Attractions that were added by a given user_id
*/
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

func GetAttractionsUnapproved() ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	var attractions []Attraction
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE approved=FALSE")
	if err != nil {
		return attractions, err
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
Gets an Attraction by parts of its title
For example: ain would find "Mainzer Backstube,Mainzer Dom,Hain ..." 
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


func GetAttractionsByCityAndType(city string,category string) ([]Attraction,error){
	var db *sql.DB = db_utils.DB
	_ = db
	var attractions []Attraction
	rows, err := db.Query("SELECT * from ATTRACTION_ENTRY WHERE type=? and approved=TRUE and city=? LIMIT 1000", city,category)
	if err != nil {
		return attractions, err
	}
	defer rows.Close()
	return getAttractionsFromDb(rows)
}