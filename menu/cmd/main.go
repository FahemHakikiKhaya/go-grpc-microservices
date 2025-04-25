package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fahemhakikikhaya/common"
	"github.com/fahemhakikikhaya/go-microservices-menu/internal/handler"
	"github.com/fahemhakikikhaya/go-microservices-menu/internal/service"
	"github.com/fahemhakikikhaya/go-microservices-menu/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:8002")
	mongoDBHost     = common.EnvString("DB_HOST", "localhost")
	mongoDBPort     = common.EnvString("DB_PORT", "27017")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	defer l.Close()
	
	// Initialize MongoDB
	db, err := initMongo()
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	menuCollection := db.Collection("menus")

	menuStore := store.NewMenuStore(menuCollection)

	menuService := service.NewMenuService(menuStore)

	handler.NewMenuHandler(grpcServer, menuService)

	
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}	
}

func initMongo() (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", mongoDBHost, mongoDBPort)
	client := common.NewMongoClient(uri)
	db := client.Database("menu-service")
	return db, nil
}
