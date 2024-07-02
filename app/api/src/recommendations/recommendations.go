package recommendations;

import(
	"src/attractions"
	"fmt"
	"src/db_utils"
	"database/sql"
)

var ErrNoRecommendation = fmt.Errorf("No Recommendations Found")

func GetRecommendationForUser(id int32,city string,pref_type string) ([]attractions.Attraction,error){
	var db *sql.DB = db_utils.DB
	var recommended_attractions []attractions.Attraction 
	rows, err := db.Query("SELECT * FROM ATTRACTION_ENTRY WHERE type = ? and city = ? ORDER BY stars LIMIT 4", pref_type,city)
	if(err != nil){
		return recommended_attractions,err
	}
	defer rows.Close()

	nodata_found := true
	var a attractions.Attraction 
	for rows.Next() {
		nodata_found = false
		rows.Scan(&a.Id,&a.Title,&a.Type,&a.Recommended_count,&a.City,&a.Info,&a.Approved,&a.PosX,&a.PosY,&a.Stars)
		recommended_attractions = append(recommended_attractions, a)

	}	
	if(nodata_found){
		return nil,ErrNoRecommendation 
	}else if(err != nil){
		return nil,err
	}

	return recommended_attractions,nil
}
