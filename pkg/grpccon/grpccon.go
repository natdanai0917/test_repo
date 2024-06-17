package grpccon

import (
	"fmt"
	"log"
	"net"

	authPb "github.com/natdanai0917/test_repo/modules/auth/authPb"
	inventoryPb "github.com/natdanai0917/test_repo/modules/inventory/inventoryPb"
	itemPb "github.com/natdanai0917/test_repo/modules/item/itemPb"
	playerPb "github.com/natdanai0917/test_repo/modules/player/playerPb"

	"github.com/natdanai0917/test_repo/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}
	grpcClientFactory struct {
		client *grpc.ClientConn
	}
	grpcAuth struct {
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.NewClient(host, opts...) //grpc.Dial is deprecated
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, fmt.Errorf("Error: client grpc client connection failed")
	}

	return &grpcClientFactory{client: clientConn}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tpc", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen:%v", err)
	}

	return grpcServer, lis
}
