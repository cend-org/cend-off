package socketio

import (
	"errors"
	"github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"log"
	"net/http"
)

var server *gosocketio.Server

func SetupSocketIOServer() {
	server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("SocketIO server starting...")
	go func() {
		log.Panic(http.ListenAndServe(":3811", serveMux))
	}()
}

func SendMessage(room string, msg any) (err error) {
	// Broadcast the message to all clients in the room
	server.BroadcastTo(room, "newMessage", msg)
	return nil
}

func JoinRoom(userId, room string) error {
	channel, err := server.GetChannel(userId)
	if err != nil {
		log.Printf("Error getting channel for client %s: %v", userId, err)
		return err
	}

	if channel == nil {
		log.Printf("Channel not found for client %s", userId)
		return errors.New("channel not found")
	}

	channel.Join(room)
	log.Printf("Client %s joined room: %s\n", userId, room)
	return nil
}

func LeaveRoom(userId, room string) error {
	channel, err := server.GetChannel(userId)
	if err != nil {
		log.Printf("Error getting channel for client %s: %v", userId, err)
		return err
	}

	if channel == nil {
		log.Printf("Channel not found for client %s", userId)
		return errors.New("channel not found")
	}

	channel.Leave(room)
	log.Printf("Client %s left room: %s\n", userId, room)
	return nil
}
