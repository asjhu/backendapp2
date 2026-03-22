package main

import (
	"database/sql"
	"log"

	"github.com/asjhu/backendapp2/api"
	db "github.com/asjhu/backendapp2/db/sqlc"
	"github.com/asjhu/backendapp2/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".") // dot mean same dir as main.go
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
