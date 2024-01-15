package sample

import "github.com/reluth/grpcserver/pb"

func NewUserFeature() *pb.UserFeature {
	numberFeature := randomInt(5, 10)
	features := []float32{}
	for i := 0; i < numberFeature; i++ {
		features = append(features, randomFloat32(5, 10))
	}
	userFeature := &pb.UserFeature{
		UserId:   randomID(),
		Features: features,
	}

	return userFeature
}
