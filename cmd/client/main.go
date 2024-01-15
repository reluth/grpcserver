package main

import (
	"context"
	"flag"
	"log"

	"github.com/reluth/grpcserver/pb"
	"github.com/reluth/grpcserver/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	userFeatureClient := pb.NewUserFeatureServiceClient(conn)

	userFeature := sample.NewUserFeature()
	reqAdd := &pb.AddUserFeatureRequest{
		UserFeature: userFeature,
	}

	resAdd, err := userFeatureClient.AddUserFeature(context.Background(), reqAdd)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("user_feature already exists")
		} else {
			log.Fatal("cannot add user_feature: ", err)
		}
	}

	log.Printf("create user_feature success, user_id: %s", resAdd.UserId)

	reqGet := &pb.GetFeatureRequest{
		UserId: resAdd.UserId,
	}
	resGet, err := userFeatureClient.GetUserFeature(context.Background(), reqGet)
	if err != nil {
		log.Fatal("cannot get user_feature: ", err)
	}

	log.Printf("get user_feature user_id: %v success, features: %v", resAdd.UserId, resGet.Feature)
}
