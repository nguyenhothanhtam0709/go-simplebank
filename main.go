package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nguyenhothanhtam0709/go-simplebank/api"
	db "github.com/nguyenhothanhtam0709/go-simplebank/db/sqlc"
	"github.com/nguyenhothanhtam0709/go-simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
