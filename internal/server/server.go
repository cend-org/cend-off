package server

import (
	"fmt"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/router"
)

func Begin() {
	defer database.CloseConnexion()

	defer recoverServer()

	router.Serve()
}

func recoverServer() {
	fmt.Print("SHUTTING DOWN")
}
