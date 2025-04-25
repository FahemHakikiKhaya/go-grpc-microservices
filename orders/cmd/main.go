package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fahemhakikikhaya/common"
	menuPb "github.com/fahemhakikikhaya/common/api/menu"
	"github.com/fahemhakikikhaya/go-microservices-orders/internal/handler"
	"github.com/fahemhakikikhaya/go-microservices-orders/internal/service"
	"github.com/fahemhakikikhaya/go-microservices-orders/internal/store"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddr        = common.EnvString("GRPC_ADDR", "localhost:8001")
	menuServiceAddr = common.EnvString("MENU_SERVICE_ADDR", "localhost:8002")
	mongoDBHost     = common.EnvString("DB_HOST", "localhost")
	mongoDBPort     = common.EnvString("DB_PORT", "27017")
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("❌ Service failed to start: %v", err)
	}
}

func run() error {
	// Connect to MenuService gRPC
	menuConn, err := grpc.Dial(menuServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to dial MenuService: %w", err)
	}
	defer menuConn.Close()
	menuClient := menuPb.NewMenuServiceClient(menuConn)

	// Initialize MongoDB
	db, err := initMongo()
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Init local gRPC server
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcAddr, err)
	}
	defer listener.Close()

	orderCollection := db.Collection("orders")

	// Init service
	store := store.NewOrderStore(orderCollection)
	service := service.NewOrderService(store)

	// Register gRPC handler
	handler.NewGRPCHandler(grpcServer, service, menuClient)

	log.Printf("🚀 gRPC server is listening on %s", grpcAddr)
	return grpcServer.Serve(listener)
}

func initMongo() (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", mongoDBHost, mongoDBPort)
	client := common.NewMongoClient(uri)
	db := client.Database("order-service")
	return db, nil
}
