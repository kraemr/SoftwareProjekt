package notifications;
import (
	_ "errors"
	"database/sql"
	"src/db_utils"
	"fmt"
)

type Notification struct{
	info string // HTML code with the message inside
	date string  // Date that notification was created on
}
var ErrNoNotification = fmt.Errorf("No Notifications Found")

func getNotificationsForId(user_id int32) ([]Notification,error){
	var db *sql.DB = db_utils.DB
	var notifications []Notification
	rows, err := db.Query("SELECT info,date FROM USER_NOTIFICATIONS WHERE user_id = ?", user_id)

	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	no_data := true	
	for rows.Next() {
		no_data = false
		notif := Notification{}
		rows.Scan(&notif.info,&notif.date)
		notifications = append(notifications,notif)
	}
	if(no_data){
		return nil,ErrNoNotification
	}
	return notifications,nil
}