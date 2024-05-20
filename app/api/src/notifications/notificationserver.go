package notifications;
import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"fmt"

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
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	_ = ws
}

func StartNotificationServer(port string,path string){
	addr := ":" + port
	http.HandleFunc(path, serveWs)
	fmt.Println("started NotificationServer at: " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

}
