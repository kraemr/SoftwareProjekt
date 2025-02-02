package reviews

import (
	"database/sql"
	_ "errors"
	"fmt"
	"src/db_utils"
)

// Review struct - representing the review data just like in the database
type Review struct {
	Id            int64   `json:id`
	Text          string  `json:text`
	Attraction_id int32   `json:attraction_id`
	Username      string  `json:username`
	User_id       int64   `json:user_id`
	Stars         float32 `json:stars`
	Date          string  `json:date`
}

// Create ErrNoReviews to return when no reviews are found
var ErrNoReviews = fmt.Errorf("No Reviews Found")

func DeleteReview(review_id int64) error {
	var db *sql.DB = db_utils.DB

	// Add prepared statement to delete the review by id
	prepared_stmt, err := db.Prepare("DELETE FROM ATTRACTION_REVIEW WHERE id=?")
	if err != nil {
		return err
	}
	// Execute the prepared statement
	result, err := prepared_stmt.Exec(review_id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func InsertReview(review Review) error {
	var db *sql.DB = db_utils.DB
	// Add prepared statement to insert a review
	prepared_stmt, err := db.Prepare("INSERT INTO ATTRACTION_REVIEW(user_id,attraction_id,text,stars,date) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	// Get the current date and use it as date for the review
	date := db_utils.GetCurrentDate()

	// Execute the prepared statement with the review data and the date
	result, err := prepared_stmt.Exec(review.User_id, review.Attraction_id, review.Text, review.Stars, date)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func UpdateReview(review Review) error {
	var db *sql.DB = db_utils.DB
	// Add prepared statement to update the review by id
	prepared_stmt, err := db.Prepare("UPDATE ATTRACTION_REVIEW SET text=?,stars=?,date=? where id = ?")
	if err != nil {
		return err
	}

	// Get the current date and use it as new date for the review
	date := db_utils.GetCurrentDate()

	// Execute the prepared statement with the review data and the date
	result, err := prepared_stmt.Exec(review.Text, review.Stars, date, review.Id)
	_ = result
	if err != nil {
		return err
	}
	return nil
}

func GetReviewsByAttractionId(attraction_id int32) ([]Review, error) {
	var db *sql.DB = db_utils.DB
	var reviews []Review

	// Get all reviews for the attraction by the attraction_id
	// Use a left join to get the username of the user who wrote the review
	rows, err := db.Query("SELECT ar.id,ar.user_id,ar.attraction_id,ar.text,ar.stars,ar.date,username FROM ATTRACTION_REVIEW as ar LEFT JOIN USER ON user_id = USER.id WHERE attraction_id = ?", attraction_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	nodata_found := true

	// Iterate over the rows and append the reviews to the list
	for rows.Next() {
		nodata_found = false
		var r Review
		rows.Scan(&r.Id, &r.User_id, &r.Attraction_id, &r.Text, &r.Stars, &r.Date, &r.Username)
		fmt.Println("found a rview")
		fmt.Println(r.Attraction_id)
		reviews = append(reviews, r)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if nodata_found {
		return nil, ErrNoReviews
	}

	return reviews, nil
}

func GetStarsForAttraction(attraction_id int32) (float32, error) {
	var db *sql.DB = db_utils.DB

	// Get all reviews for the attraction by the attraction_id
	// Limit the query to 1000 reviews to prevent long loading times
	rows, err := db.Query("SELECT * FROM ATTRACTION_REVIEW WHERE attraction_id = ? LIMIT 1000", attraction_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer rows.Close()
	var r Review

	star_sum := float32(0)
	reviews_count := int32(0)

	// Iterate over the rows and calculate the sum of the stars and the count of the reviews
	for rows.Next() {
		rows.Scan(&r.Id, &r.User_id, &r.Attraction_id, &r.Text, &r.Stars, &r.Date)
		fmt.Println(r)
		star_sum += r.Stars
		reviews_count += 1
	}

	if reviews_count == 0 {
		return 0, ErrNoReviews
	}

	return star_sum / float32(reviews_count), nil

}

func GetReviewsByUserId(user_id int32) ([]Review, error) {
	var db *sql.DB = db_utils.DB
	var reviews []Review

	// Get all reviews for the user by the user_id
	// Use a left join to get the username of the user who wrote the review
	rows, err := db.Query("SELECT ar.id,ar.user_id,ar.attraction_id,ar.text,ar.stars,ar.date,username FROM ATTRACTION_REVIEW as ar LEFT JOIN USER ON user_id = USER.id WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	nodata_found := true

	// Iterate over the rows and append the reviews to the list
	for rows.Next() {
		nodata_found = false
		r := Review{}
		rows.Scan(&r.Id, &r.User_id, &r.Attraction_id, &r.Text, &r.Stars, &r.Date, &r.Username)
		reviews = append(reviews, r)
	}
	if err != nil {
		return nil, err
	}
	if nodata_found {
		return nil, ErrNoReviews
	}
	return reviews, nil
}
