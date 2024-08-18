package main

import (
	"context"
	"log"
	"os"

	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/pkg/database"
	"github.com/natdanai0917/test_repo/server"
)

func main() {
	ctx := context.Background()

	//initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error : .env path is required")
		}
		return os.Args[1]
	}())

	//Database Connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	//Start Server
	server.Start(ctx, &cfg, db)
}
