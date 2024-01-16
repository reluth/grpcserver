package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/reluth/grpcserver/model"
	"github.com/reluth/grpcserver/pb"
)

var (
	ErrAlreadyExists = errors.New("already exists")
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
	err := store.db.Create(model.NewUserFeature(userFeature.GetUserId(), userFeature.GetFeatures())).Error

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (store *DBUserFeatureStore) Find(userID string) (*pb.UserFeature, error) {
	var res model.UserFeature

	err := store.db.Where(&model.UserFeature{UserID: userID}).Take(&res).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAlreadyExists
		}
		return nil, err
	}

	return &pb.UserFeature{
		UserId:   userID,
		Features: res.Features,
	}, err
}
