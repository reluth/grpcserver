package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/reluth/grpcserver/database"
	"github.com/reluth/grpcserver/model"
	"github.com/reluth/grpcserver/pb"
	"github.com/reluth/grpcserver/service"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db := database.Init()

	err := db.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.UserFeature{})

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("start server on port %d", port)

	userFeatureServer := service.NewUserFeatureServer(service.NewDBUserFeatureStore(db))
	grpcServer := grpc.NewServer()
	pb.RegisterUserFeatureServiceServer(grpcServer, userFeatureServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
