package service

import (
	"encoding/json"

	"github.com/ragsharan/foss-apis/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type profileService struct{}

type IProfileService interface {
	GetProfile(data []byte) (primitive.M, error)
	AddProfile(data []byte) (*mongo.InsertOneResult, error)
	UpdateProfile(data []byte) (*mongo.UpdateResult, error)
	RemoveProfile(data []byte) (*mongo.DeleteResult, error)
}

func InstProfileService() IProfileService {
	return &profileService{}
}

var (
	repoMongo repository.IMongoRepository = repository.ObjIMongoRepository()
)

const collection string = "profile"

func (*profileService) GetProfile(data []byte) (primitive.M, error) {
	kmap := make(map[string]interface{})
	json.Unmarshal(data, &kmap)
	result, err := repoMongo.FindOne(collection, kmap)
	return result, err
}
func (*profileService) AddProfile(data []byte) (*mongo.InsertOneResult, error) {
	kmap := make(map[string]interface{})
	json.Unmarshal(data, &kmap)
	result, err := repoMongo.Create(collection, kmap)
	return result, err
}
func (*profileService) UpdateProfile(data []byte) (*mongo.UpdateResult, error) {
	kmap := make(map[string]interface{})
	json.Unmarshal(data, &kmap)
	result, err := repoMongo.UpdateOne(collection, kmap)
	return result, err
}
func (*profileService) RemoveProfile(data []byte) (*mongo.DeleteResult, error) {
	kmap := make(map[string]interface{})
	json.Unmarshal(data, &kmap)
	result, err := repoMongo.DeleteDocument(collection, kmap)
	return result, err
}
