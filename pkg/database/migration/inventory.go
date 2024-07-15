package migration

import (
	"context"
	"log"

	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func inventoryDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("inventory_db")
}

func InventoryMigrate(pctx context.Context, cfg *config.Config) {
	db := inventoryDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	//inventory
	col := db.Collection("players_inventory")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"player_id", 1}, {"item_id", 1}}},
	})
	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	//roles
	col = db.Collection("players_inventory_queue")

	results, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate inventory completed: ", results)
}
