package main

import (
	"duval/internal/route"
	"duval/pkg/database"
	"github.com/jmoiron/sqlx"
)

func main() {
	defer func(Client *sqlx.DB) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(database.Client)

	err := route.Serve()
	if err != nil {
		panic(err)
	}
}
