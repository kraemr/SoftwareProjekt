package favorites;

import(
	"src/db_utils"
	"fmt"
	"database/sql"
)

type AttractionFavorite struct{
	Id int64 `json:id`
	User_id int32 `json:user_id`
	Attraction_id int32 `json:attraction_id`
//	Type string `json:type`
//	City string `json:city`
};

type AttractionFavoriteUnion struct{
	Id 				  int64 	`json:id`
	User_id 		  int32 	`json:user_id`
	Attraction_id     int64	    `json:attraction_id`
	Title 			  string    `json:"title"`
	Type  			  string    `json:"type"`
	Recommended_count int 	    `json:"recommended_count"`
	City 			  string    `json:"city"`
	Info 			  string    `json:"info"`
	Approved		  bool	    `json:approved`
	PosX 			  float32   `json:"posX"`
	PosY 			  float32   `json:"posY"`
	Stars			  float32   `json:stars`
};



// IF this was returned in recommendations then we dont send any
var ErrNoFavorites = fmt.Errorf("No Favorites Found")

/*
CREATE TABLE USER_FAVORITE(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    attraction_id INT NOT NULL,
    type varchar(32) NOT NULL,
    city varchar(32) NOT NULL
);
*/


func DeleteAttractionFavoriteById(id int32) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("DELETE FROM USER_FAVORITE WHERE id=?");
	if(err != nil){
		fmt.Println("Couldnt Remove Attraction Favorite")
		return err
	}
	result,err1 := prepared_stmt.Exec(id)
	_ = result
	if err1 != nil {
		return err1
	}
	return nil
}


func DeleteAttractionFavoriteByAttractionId(attraction_id int32) error {
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("DELETE FROM USER_FAVORITE WHERE id=?");
	if(err != nil){
		fmt.Println("Couldnt Remove Attraction Favorite")
		return err
	}
	result,err1 := prepared_stmt.Exec(attraction_id)
	if err1 != nil {
		return err1
	}	
	rows_affected,row_err := result.RowsAffected()
	if(row_err != nil || rows_affected == 0){
		return row_err
	}

	return nil
}


// TODO: Check for IDOR, Check for proper Auth here
func AddAttractionFavoriteById(user_id int32,attraction_id int32) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("INSERT INTO USER_FAVORITE(user_id,attraction_id) VALUES(?,?)");
	if(err != nil){
		fmt.Println("Couldnt Add Attraction Favorite")
		return err
	}
	result,err1 := prepared_stmt.Exec(user_id,attraction_id)
	_ = result
	if(err1 != nil){
		return err1
	}
	return nil
}




/*
JOIN User_favorites and the corresponding Attraction
*/
// TODO: Check for IDOR, Check for proper Auth here
func GetAttractionFavoritesByUserId(user_id int32) ([]AttractionFavoriteUnion,error){
	var db *sql.DB = db_utils.DB
	var favorites []AttractionFavoriteUnion
	rows,err := db.Query("SELECT * FROM USER_FAVORITE as uf JOIN ATTRACTION_ENTRY as at ON uf.attraction_id = at.id WHERE user_id=?",user_id)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	nodata_found := true
	fmt.Println(user_id)
	for rows.Next() {
		var fav AttractionFavoriteUnion = AttractionFavoriteUnion{};
		rows.Scan(
			&fav.Id,
			&fav.User_id,
			&fav.Attraction_id,
			&fav.Attraction_id,
			&fav.Title,
			&fav.Type,
			&fav.Recommended_count,
			&fav.City,
			&fav.Info,
			&fav.Approved,
			&fav.PosX,
			&fav.PosY,
			&fav.Stars,
		)
		fmt.Println(fav)
		nodata_found = false
		favorites = append(favorites,fav)
	}

	if(err != nil){
		return nil,err
	}else if(nodata_found){
		return nil,ErrNoFavorites
	}else{
		return favorites,nil
	}
}