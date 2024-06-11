package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/cend-org/duval/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tanimutomo/sqlfile"
	"log"
)

type Config struct {
	Port                    string `toml:"port"`
	Host                    string `toml:"host"`
	TokenSecret             string `toml:"token_secret"`
	DatabaseUserName        string `toml:"database_user_name"`
	DatabaseUserPassword    string `toml:"database_user_password"`
	DatabaseName            string `toml:"database_name"`
	DatabaseHost            string `toml:"database_host"`
	DatabasePort            string `toml:"database_port"`
	DatabaseConnexionString string
	Mode                    string `toml:"mode"`
}

func main() {

	var config Config
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DatabaseUserName, config.DatabaseUserPassword,
		config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	fmt.Printf("Opening Connection ... %s\n", databaseUrl)

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	s := sqlfile.New()

	migrationDir := fmt.Sprintf("%s/db/migrator.sql", utils.MIGRATION_DIR)
	fmt.Printf("Running migration ... %s\n", migrationDir)

	if err := s.File(migrationDir); err != nil {
		log.Fatalf("Error reading migration file: %v", err)
	}

	res, err := s.Exec(db)
	if err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}

	for _, r := range res {
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			log.Fatalf("Error fetching rows affected: %v", err)
		}
		fmt.Printf("Migration successful, rows affected: %d\n", rowsAffected)
	}
}
