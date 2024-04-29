package main

import (
	"github.com/cend-org/duval/internal/server"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	server.Begin()
}
