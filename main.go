package main

import (
	"context"
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/pkg/database"
	"github.com/natdanai0917/test_repo/server"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	_ = ctx

	//initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error : .env path is required")
		}
		return os.Args[1]
	}())

	//Database Connection
	db := database.DbConn(ctx, &cfg)
	log.Println(db)

	//Start Server
	server.Start(ctx, &cfg, db)
}
