//go:generate go run cmd/generate/generate.go

package main

import (
	"fmt"
	"github.com/cend-org/duval/internal/server"
)

import _ "github.com/go-sql-driver/mysql"

const version string = "0.0.2 - dev"

func main() {
	fmt.Println("\nApp running version ", version)

	server.Begin()
}
