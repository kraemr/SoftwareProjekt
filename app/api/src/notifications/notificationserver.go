package notifications

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"src/sessions"
	"src/users"
	"time"

	"github.com/gorilla/websocket"
)	

var(
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)


/*
Currently this does not Check the origin
*/
func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	_ = ws
}

// User sends Id to start sesh
// c.readmessage
// parse json -> jsonObject
// getNotificationsForID(jsonObject.Id) -> send Notifications
var NotificationSendSignal bool = false


/*
This sends Notifications to users connected to the websocket
every 10 seconds it checks if there are notifications and sends them
*/
func sendNotifications(w http.ResponseWriter, r *http.Request){
	if(!sessions.CheckLoggedIn(r)) {
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// check the users date in json and only send those
	// unless the user sends the all:true flag
	// otherwise only NEW Notifications will be sent to client
	id := sessions.GetLoggedInUserId(r)
	fmt.Println(id)
	// we have user id, now get the city
	city,err := users.GetUserCityById(id)
	if(err != nil){
		fmt.Println("user has no city???")
		return
	}
	fmt.Println(city);
	for {
			time.Sleep(10 * time.Second)
			
			user_notifications,user_err := getNotificationsForId(id)
			if(user_err != nil){
				fmt.Println(err.Error())
				break
			}
			city_notifications,city_err := getRecentNotificationsForCity(city)
			if(city_err != nil){
				fmt.Println(err.Error())
				break
			}
			notifications := append(user_notifications,city_notifications...)

			json_bytes , json_err := json.Marshal(notifications)
			if(json_err != nil){
				fmt.Println("error Creating Json")
				break
			}
			err := c.WriteMessage(websocket.TextMessage, json_bytes)
			if err != nil {
				log.Println("write:", err)
				break
			}
		
	}
}

// Starts the Websocket
func StartNotificationServer(port string,path string){
	addr := ":" + port
	http.HandleFunc(path, sendNotifications)
	fmt.Println("started NotificationServer at: " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

}
