package playerRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/natdanai0917/test_repo/modules/player"
	"github.com/natdanai0917/test_repo/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PlayerRepositoryService interface {
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error
		GetPlayerSavingAccount(ptcx context.Context, playerId string) (*player.PlayerSavingAccount, error)
		FindOnePlayerCredential(pctx context.Context, email string) (*player.Player, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*player.Player, error) 
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playerRepository{db}
}

func (r *playerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}

func (r *playerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	player := new(player.Player)
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": username},
			{"email": email},
		}},
	).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}
	return false
}

func (r *playerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	playerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player failed")
	}

	return playerId.InsertedID.(primitive.ObjectID), nil
}

func (u *playerRepository) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := u.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.PlayerProfileBson)

	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjId(playerId)},
		options.FindOne().SetProjection(
			bson.M{
				"_id":        1,
				"email":      1,
				"username":   1,
				"created_at": 1,
				"updated_at": 1,
			},
		),
	).Decode(result); err != nil {
		log.Printf("Error : FineOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error: player profile not found")
	}

	return result, nil
}

func (r *playerRepository) InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Result: InsertOnePlayerTransaction: %v", result.InsertedID)
	}

	return nil

}

func (r *playerRepository) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error) {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	filter :=
		bson.A{
			bson.D{{"$match", bson.D{{"player_id", playerId}}}},
			bson.D{
				{"$group",
					bson.D{
						{"_id", "$player_id"},
						{"balance", bson.D{{"$sum", "$amount"}}},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"player_id", "$_id"},
						{"_id", 0},
						{"balance", 1},
					},
				},
			},
		}

	cursors, err := col.Aggregate(ctx, filter)
	if err != nil {
		log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
		return nil, errors.New("error: failed to get player saving account")
	}

	result := new(player.PlayerSavingAccount)
	for cursors.Next(ctx) {
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
			return nil, errors.New("error: failed to get player saving account")
		}
	}

	return result, nil
}

func (r *playerRepository) FindOnePlayerCredential(pctx context.Context, email string) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.Player)

	if err := col.FindOne(ctx, bson.M{"email": email}).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerCredential: %s", err.Error())
		return nil, errors.New("error: email is invalid")
	}
	return result, nil
}

func (r *playerRepository) FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.Player)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjId(playerId)}).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerProfileToRefresh: %s", err.Error())
		return nil, errors.New("error: player profile not found")
	}

	return result, nil
}