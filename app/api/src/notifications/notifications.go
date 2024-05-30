package notifications;
import (
	_ "errors"
	"database/sql"
	"src/db_utils"
	"fmt"
)

type Notification struct{
	Info string `json:info` // HTML code with the message inside
	Date string `json:date`// Date that notification was created on
}

var ErrNoNotification = fmt.Errorf("No Notifications Found")


func AddNotification(user_id int32,notification Notification) error{
	return nil
}

func getNotificationsFromDb(rows *sql.Rows)  ([]Notification,error){
	var notifications []Notification
	no_data := true	
	for rows.Next() {
		no_data = false
		notif := Notification{}
		rows.Scan(&notif.Info,&notif.Date)
		notifications = append(notifications,notif)
	}
	if(no_data){
		return nil,ErrNoNotification
	}
	return notifications,nil
}

/*
If user was not logged in when notification came
this can be used to get the notifications
*/

func getNotificationsForIDByDate(user_id int32,date string) ([]Notification,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT info,date FROM USER_NOTIFICATIONS WHERE user_id = ? and date > ? ", user_id,date)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	return getNotificationsFromDb(rows);
}

func getNotificationsForId(user_id int32) ([]Notification,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT info,date FROM USER_NOTIFICATIONS WHERE user_id = ? LIMIT 100 ORDER BY date DESC", user_id)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	return getNotificationsFromDb(rows);
}

func getRecentNotificationsForCity(city string) ([]Notification,error){
	var db *sql.DB = db_utils.DB
	rows, err := db.Query("SELECT info,date FROM CITY_NOTIFICATIONS WHERE city = ? LIMIT 100 ORDER BY date DESC", city)
	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	return getNotificationsFromDb(rows);
}