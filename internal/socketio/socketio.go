package socketio

import (
	"github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"net/http"
)

var socketIoServer *gosocketio.Server

func SetupSocketIOServer() {
	socketIoServer = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	socketIoServer.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
	})

	socketIoServer.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("client Disconnected")
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", socketIoServer)
	log.Panic(http.ListenAndServe(":8086", serveMux))

}

func SendMessage(room string, msg any) (err error) {
	err = socketIoServer.On("message", func(c *gosocketio.Channel) string {
		//send event to all in room
		c.BroadcastTo(room, "message", msg)
		return "OK"
	})
	if err != nil {
		return err
	}
	return
}
