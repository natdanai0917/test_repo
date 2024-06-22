package migration

import (
	"context"
	"log"

	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/player"
	"github.com/natdanai0917/test_repo/pkg/database"
	"github.com/natdanai0917/test_repo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func playerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config) {
	db := playerDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("player_transactions")

	//indexes
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"player_id", 1}}},
	})
	log.Println(indexs)

	col = db.Collection("players")

	//indexes
	//indexes
	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"email", 1}}},
	})
	log.Println(indexs)

	//roles data
	documents := func() []any {
		roles := []*player.Player{
			{
				Email:    "player001@sekai.com",
				Password: "123456",
				Username: "player001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player002@sekai.com",
				Password: "123456",
				Username: "player002",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player003@sekai.com",
				Password: "123456",
				Username: "player003",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "admin001@sekai.com",
				Password: "123456",
				Username: "admin001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}
		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player completed: ", results)

	playerTransactions := make([]any, 0)
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, &player.PlayerTransaction{
			PlayerId:  "player" + p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime(),
		})
	}

	col = db.Collection("player_transactions")
	results, err = col.InsertMany(pctx, playerTransactions, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player_transactions completed", results)

	col = db.Collection("player_transaction_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player_transactions_queue completed: ", result)
}
