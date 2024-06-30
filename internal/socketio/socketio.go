package socketio

import (
	"github.com/rs/cors"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"net/http"
)

var (
	server *socket.Server
	client *socket.Socket
)

func SetupSocketIOServer() {

	server = socket.NewServer(nil, nil)

	server.On("connection", func(clients ...any) {
		client = clients[0].(*socket.Socket)

		client.On("newMessage", func(datas ...any) {
			if len(datas) > 0 {
				client.Emit("newMessage", datas)
			}
		})

		client.On("joinRoom", func(datas ...any) {
			if len(datas) > 0 {
				if room, ok := datas[0].(string); ok {
					clientRoom := socket.Room(room)
					client.Join(clientRoom)
				}
			}
		})

		client.On("leaveRoom", func(datas ...any) {
			if len(datas) > 0 {
				if room, ok := datas[0].(string); ok {
					client.Leave(socket.Room(room))
				}
			}
		})
	})

	server.On("disconnect", func(clients ...any) {
		log.Println("Client disconnected")
	})

	corsHandler := cors.Default().Handler(server.ServeHandler(nil))

	http.Handle("/socket.io/", corsHandler)

	log.Println("SocketIO server starting...")
	go func() {
		log.Panic(http.ListenAndServe(":8086", nil))
	}()
}

func SendMessage(message any) (err error) {
	//err = client.To(socket.Room(receiverId)).To(socket.Room(userId)).Emit("private message", message)
	err = client.Emit("newMessage", message)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return err
	}
	return nil
}
