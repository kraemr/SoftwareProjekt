package notifications;
import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)	

func BroadcastRecommendations(){

}

var(
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)
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

func sendNotifications(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// check the users date in json and only send those
	// unless the user sends the all:true flag
	// otherwise only NEW Notifications will be sent to clients
	for {
		notifications,err := getNotificationsForId(911111)
		if(err != nil){
			fmt.Println(err.Error())
		}
		json_bytes , json_err := json.Marshal(notifications)
		if(json_err != nil){
			fmt.Println("error Creating Json")
		}
		err = c.WriteMessage(websocket.TextMessage, json_bytes)
		if err != nil {
			log.Println("write:", err)
			//break
		}
		break
	}
}




func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Println(message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}

func StartNotificationServer(port string,path string){
	addr := ":" + port
	http.HandleFunc(path, sendNotifications)
	fmt.Println("started NotificationServer at: " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

}
