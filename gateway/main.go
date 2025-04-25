package main

import (
	"log"
	"net/http"

	"github.com/fahemhakikikhaya/common"
	menuPb "github.com/fahemhakikikhaya/common/api/menu"
	orderPb "github.com/fahemhakikikhaya/common/api/order"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8000")
	orderServiceAddr = "localhost:8001"
	menuServiceAddr = "localhost:8002"
)

func main() {
	orderConn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()) )

	if err != nil {
		log.Fatal("Failed to dial server %v", err)
	}

	defer orderConn.Close()

	menuConn, err :=  grpc.Dial(menuServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()) )

	if err != nil {
		log.Fatal("Failed to dial server %v", err)
	}

	defer menuConn.Close()

	orderService := orderPb.NewOrderServiceClient(orderConn)
	menuService := menuPb.NewMenuServiceClient(menuConn)



	log.Println("Dialing orders service at", orderServiceAddr)
	log.Println("Dialing menu service at", menuServiceAddr)


	mux := http.NewServeMux()
	handler := NewHandler(orderService, menuService)
	handler.registerRoutes(mux)

	log.Printf("Starting http server on port this: %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}