package server

import (
	"fmt"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/router"
	"github.com/cend-org/duval/internal/socketio"
)

func Begin() {
	defer database.CloseConnexion()

	defer recoverServer()

	socketio.SetupSocketIOServer()

	router.Serve()
}

func recoverServer() {
	fmt.Print("SHUTTING DOWN")
}
