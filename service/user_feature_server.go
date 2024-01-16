package service

import (
	"context"
	"errors"
	"log"

	"github.com/reluth/grpcserver/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserFeatureServer struct {
	pb.UserFeatureServiceServer
	userFeatureStore UserFeatureStore
}

func NewUserFeatureServer(userFeatureStore UserFeatureStore) *UserFeatureServer {
	return &UserFeatureServer{
		userFeatureStore: userFeatureStore,
	}
}

func (server *UserFeatureServer) AddUserFeature(ctx context.Context, req *pb.AddUserFeatureRequest) (*pb.AddUserFeatureResponse, error) {
	userFeature := req.GetUserFeature()
	userID := userFeature.GetUserId()
	log.Printf("receive a add user_feature request with user_id: %s", userID)

	err := server.userFeatureStore.Save(userFeature)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot add user_feature to the store: %v", err)
	}

	log.Printf("add user_feature with user_id: %s", userID)

	res := &pb.AddUserFeatureResponse{
		UserId: userID,
	}
	return res, nil
}

func (server *UserFeatureServer) GetUserFeature(ctx context.Context, req *pb.GetFeatureRequest) (*pb.GetFeatureResponse, error) {
	userID := req.GetUserId()
	log.Printf("receive a get user_feature request with user_id: %s", userID)

	userFeature, err := server.userFeatureStore.Find(userID)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.NotFound
		}
		return nil, status.Errorf(code, "cannot get user_feature with user_id: %v", err)
	}

	log.Printf("get user_feature with user_id: %s", userID)

	feature := []float32{}
	if userFeature != nil {
		feature = userFeature.Features
	}

	res := &pb.GetFeatureResponse{
		Feature: feature,
	}
	return res, nil
}
