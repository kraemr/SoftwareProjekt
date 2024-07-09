package favorites

import (
	"database/sql"
	"fmt"
	"src/db_utils"
)

// AttrationFavorite and AttractionFavoriteUnion structs - representing the favorite data just like in the database
type AttractionFavorite struct {
	Id            int64 `json:id`
	User_id       int64 `json:user_id`
	Attraction_id int64 `json:attraction_id`
}

type AttractionFavoriteUnion struct {
	Id                int64   `json:id`
	User_id           int64   `json:user_id`
	Attraction_id     int64   `json:attraction_id`
	Title             string  `json:"title"`
	Type              string  `json:"type"`
	Recommended_count int     `json:"recommended_count"`
	City              string  `json:"city"`
	Street            string  `json:street`
	Housenumber       string  `json:housenumber`
	Info              string  `json:"info"`
	Approved          bool    `json:approved`
	PosX              float32 `json:"posX"`
	PosY              float32 `json:"posY"`
	Stars             float32 `json:stars`
	Img_url           string  `json:img_url`
	Added_by          string  `json:added_by`
}

// Create ErrNoFavorites to return when no favorites are found
var ErrNoFavorites = fmt.Errorf("No Favorites Found")

func DeleteAttractionFavoriteById(id int64) error {
	var db *sql.DB = db_utils.DB

	// Add prepared statement to delete the favorite by id
	prepared_stmt, err := db.Prepare("DELETE FROM USER_FAVORITE WHERE id=?")
	if err != nil {
		fmt.Println("Couldnt Remove Attraction Favorite")
		return err
	}

	// Execute the prepared statement
	result, err1 := prepared_stmt.Exec(id)
	_ = result

	if err1 != nil {
		return err1
	}
	return nil
}

func DeleteAttractionFavoriteByAttractionId(attraction_id int64, user_id int64) error {
	var db *sql.DB = db_utils.DB

	// Add prepared statement to delete the favorite by id and user_id
	prepared_stmt, err := db.Prepare("DELETE FROM USER_FAVORITE WHERE user_id=? and attraction_id=?")

	// Error handling: return error if prepared statement could not be created
	if err != nil {
		fmt.Println("Couldnt Remove Attraction Favorite")
		return err
	}

	// Execute the prepared statement
	result, err1 := prepared_stmt.Exec(user_id, attraction_id)
	if err1 != nil {
		return err1
	}

	// Error handling: return error if no rows were affected => favorite could not be deleted
	rows_affected, row_err := result.RowsAffected()
	if row_err != nil || rows_affected == 0 {
		return row_err
	}

	return nil
}

func CheckFavoriteExists(attraction_id int64, user_id int64) (bool, error) {
	var db *sql.DB = db_utils.DB

	rows, err := db.Query("SELECT COUNT(*) FROM USER_FAVORITE WHERE attraction_id = ? and user_id = ?", attraction_id, user_id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var count int64
	rows.Next() // iterate over the rows
	rows.Scan(&count)

	// Check if the count is 0, if so return false => no favorite exists
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func GetAttractionFavoriteCountByAttractionId(attraction_id int64) (int, error) {
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT COUNT(*) FROM USER_FAVORITE WHERE attraction_id = ?", attraction_id)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Check if any data is available
	nodata_found := true
	var count int
	for rows.Next() {
		rows.Scan(
			&count,
		)
		nodata_found = false
	}

	// Error handling: return error if no data was found or something went wrong while scanning, otherwise return the count
	if err != nil {
		return 0, err
	} else if nodata_found {
		return 0, ErrNoFavorites
	} else {
		return count, nil
	}
}

// TODO: Check for IDOR, Check for proper Auth here
func AddAttractionFavoriteById(user_id int64, attraction_id int64) error {
	var db *sql.DB = db_utils.DB

	// Add prepared statement to insert a new user favorite
	prepared_stmt, err := db.Prepare("INSERT INTO USER_FAVORITE(user_id,attraction_id) VALUES(?,?)")
	if err != nil {
		fmt.Println("Couldnt Add Attraction Favorite")
		return err
	}
	result, err1 := prepared_stmt.Exec(user_id, attraction_id)
	_ = result
	if err1 != nil {
		return err1
	}
	return nil
}

func GetAttractionFavoritesByAttractionId(attraction_id int64) ([]AttractionFavoriteUnion, error) {
	var db *sql.DB = db_utils.DB
	var fav_list []AttractionFavoriteUnion

	// Construct the query to get every favorite by attraction_id
	// Join the USER_FAVORITE and ATTRACTION_ENTRY tables to get the favorites and the corresponding attractions
	rows, err := db.Query("SELECT * FROM USER_FAVORITE as uf JOIN ATTRACTION_ENTRY as at ON uf.attraction_id = at.id WHERE attraction_id=?", attraction_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	nodata_found := true

	// Iterate over the rows and append the favorites to the list
	for rows.Next() {
		var fav AttractionFavoriteUnion = AttractionFavoriteUnion{}
		rows.Scan(
			&fav.Id,
			&fav.User_id,
			&fav.Attraction_id,
			&fav.Attraction_id,
			&fav.Title,
			&fav.Type,
			&fav.Recommended_count,
			&fav.City,
			&fav.Street,
			&fav.Housenumber,
			&fav.Info,
			&fav.Approved,
			&fav.PosX,
			&fav.PosY,
			&fav.Stars,
			&fav.Img_url,
			&fav.Added_by,
		)
		nodata_found = false
		fav_list = append(fav_list, fav)
	}

	// Error handling: return error if no data was found or something went wrong while scanning, otherwise return the list of favorites
	if err != nil {
		return nil, err
	} else if nodata_found {
		return nil, ErrNoFavorites
	} else {
		return fav_list, nil
	}
}

/*
JOIN User_favorites and the corresponding Attraction
*/
// TODO: Check for IDOR, Check for proper Auth here
func GetAttractionFavoritesByUserId(user_id int64) ([]AttractionFavoriteUnion, error) {
	var db *sql.DB = db_utils.DB
	var fav_list []AttractionFavoriteUnion

	// Construct the query to get every favorite by user_id
	// Join the USER_FAVORITE and ATTRACTION_ENTRY tables to get the favorites and the corresponding attractions
	rows, err := db.Query("SELECT * FROM USER_FAVORITE as uf JOIN ATTRACTION_ENTRY as at ON uf.attraction_id = at.id WHERE user_id=?", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	nodata_found := true

	// Iterate over the rows and append the favorites to the list
	for rows.Next() {
		var fav AttractionFavoriteUnion = AttractionFavoriteUnion{}
		rows.Scan(
			&fav.Id,
			&fav.User_id,
			&fav.Attraction_id,
			&fav.Attraction_id,
			&fav.Title,
			&fav.Type,
			&fav.Recommended_count,
			&fav.City,
			&fav.Street,
			&fav.Housenumber,
			&fav.Info,
			&fav.Approved,
			&fav.PosX,
			&fav.PosY,
			&fav.Stars,
			&fav.Img_url,
			&fav.Added_by,
		)
		nodata_found = false
		fav_list = append(fav_list, fav)
	}

	if err != nil {
		return nil, err
	} else if nodata_found {
		return nil, ErrNoFavorites
	} else {
		return fav_list, nil
	}
}
