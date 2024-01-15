package service

import (
	"github.com/jinzhu/gorm"
	"github.com/reluth/grpcserver/model"
	"github.com/reluth/grpcserver/pb"
)

type UserFeatureStore interface {
	Save(userFeature *pb.UserFeature) error
	Find(userID string) (*pb.UserFeature, error)
}

type DBUserFeatureStore struct {
	db *gorm.DB
}

func NewDBUserFeatureStore(db *gorm.DB) *DBUserFeatureStore {
	return &DBUserFeatureStore{db: db}
}

func (store *DBUserFeatureStore) Save(userFeature *pb.UserFeature) error {
	err := store.db.Save(model.NewUserFeature(userFeature.GetUserId(), userFeature.GetFeatures())).Error

	if err != nil {
		return err
	}

	return nil
}

func (store *DBUserFeatureStore) Find(userID string) (*pb.UserFeature, error) {
	var res model.UserFeature

	err := store.db.Where(&model.UserFeature{UserID: userID}).Take(&res).Error

	if err != nil {
		return nil, err
	}

	return &pb.UserFeature{
		UserId:   userID,
		Features: res.Features,
	}, err
}
