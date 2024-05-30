package reviews;
import (
	_ "errors"
	"database/sql"
	"src/db_utils"
	"fmt"
)

type Review struct{
	Id int64			    `json:id`
	Text string     	    `json:text`
	Attraction_id int32     `json:attraction_id`
	Username string 	    `json:username`
	User_id int64   	    `json:user_id` 
	Stars float32   		`json:stars`
	Date string				`json:date`
}
var ErrNoReviews = fmt.Errorf("No Reviews Found")

func DeleteReview(review_id int64) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("DELETE FROM ATTRACTION_REVIEW WHERE id=?")
	if(err != nil){
		return err
	}
	result,err := prepared_stmt.Exec(review_id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}



func InsertReview(review Review) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("INSERT INTO ATTRACTION_REVIEW(user_id,attraction_id,text,stars,date) VALUES(?,?,?,?,?)")
	if(err != nil){
		return err
	}

	date := db_utils.GetCurrentDate()
	result,err := prepared_stmt.Exec(review.User_id,review.Attraction_id,review.Text,review.Stars,date)
	_ = result
	if err != nil {
		return err
	}
	return nil
}


func UpdateReview(review Review) error{
	var db *sql.DB = db_utils.DB
	prepared_stmt,err := db.Prepare("UPDATE ATTRACTION_REVIEW SET text=?,stars=?,date=? where id = ?")
	if(err != nil){
		return err
	}

	date := db_utils.GetCurrentDate()
	result,err := prepared_stmt.Exec(review.Text,review.Stars,date,review.Id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func GetReviewsByAttractionId(attraction_id int32) ([]Review,error){
	var db *sql.DB = db_utils.DB
	var reviews []Review
	rows, err := db.Query("SELECT * FROM ATTRACTION_REVIEW WHERE attraction_id = ?",attraction_id)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	nodata_found := true

	for rows.Next(){
		nodata_found = false
		var r Review

		rows.Scan(&r.Id,&r.User_id,&r.Attraction_id,&r.Text,&r.Stars,&r.Date)
		fmt.Println("found a rview")
		fmt.Println(r.Attraction_id)
		reviews = append(reviews,r)
	}
	if(err != nil){	
		return nil,err
	}
	if(nodata_found){
		return nil,ErrNoReviews
	}

	return reviews,nil
}

func GetReviewsByUserId(user_id int32) ([]Review,error){
	var db *sql.DB = db_utils.DB
	var reviews []Review
	rows, err := db.Query("SELECT * FROM ATTRACTION_REVIEW WHERE user_id = ?",user_id)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	nodata_found := true

	for rows.Next(){
		nodata_found = false
		r := Review{}
		rows.Scan(&r.Id,&r.User_id,&r.Attraction_id,&r.Text,&r.Stars)
		reviews = append(reviews,r)
	}
	if(err != nil){	
		return nil,err
	}
	if(nodata_found){
		return nil,ErrNoReviews
	}
	return reviews,nil
}


