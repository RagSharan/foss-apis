package service

import (
	"encoding/json"
	"log"

	"github.com/ragsharan/foss-apis/entity"
	"github.com/ragsharan/foss-apis/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type postService struct{}

type IPostService interface {
	FindPost(data []byte) (primitive.M, error)
	FindPostList(data []byte) ([]primitive.M, error)
	CreatePost(data []byte) (*mongo.InsertOneResult, error)
	CreatePostList(data []byte) (*mongo.InsertManyResult, error)
	UpdatePost(data []byte) (*mongo.UpdateResult, error)
	UpdatePostList(data []byte) (*mongo.UpdateResult, error)
	DeletePost(data []byte) (*mongo.DeleteResult, error)
}

func InstPostService() IPostService {
	return &postService{}
}

const CollectionName string = "post"

var (
	mongoRepo repository.IMongoRepository = repository.ObjIMongoRepository()
)

func (*postService) FindPost(data []byte) (primitive.M, error) {
	temp := string(data)
	kmap := make(map[string]interface{})
	json.Unmarshal([]byte(temp), &kmap)
	log.Println("postService:", kmap)
	result, err := mongoRepo.FindOne(CollectionName, kmap)
	return result, err
}

func (*postService) FindPostList(data []byte) ([]primitive.M, error) {
	temp := string(data)
	kmap := make(map[string]interface{})
	json.Unmarshal([]byte(temp), &kmap)
	result, err := mongoRepo.FindList(CollectionName, kmap)
	return result, err
}

func (*postService) CreatePost(data []byte) (*mongo.InsertOneResult, error) {
	var post entity.Post
	json.Unmarshal(data, &post)
	log.Println("create post", post)
	result, err := mongoRepo.Create(CollectionName, post)
	return result, err
}
func (*postService) CreatePostList(data []byte) (*mongo.InsertManyResult, error) {
	var post []interface{}
	json.Unmarshal(data, &post)
	result, err := mongoRepo.CreateMany(CollectionName, post)

	return result, err
}
func (*postService) UpdatePost(data []byte) (*mongo.UpdateResult, error) {
	temp := string(data)
	kmap := make(map[string]interface{})
	json.Unmarshal([]byte(temp), &kmap)
	result, err := mongoRepo.UpdateOne(CollectionName, kmap)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func (*postService) UpdatePostList(data []byte) (*mongo.UpdateResult, error) {
	temp := string(data)
	kmap := make(map[string]interface{})
	json.Unmarshal([]byte(temp), &kmap)
	result, err := mongoRepo.UpdateMany(CollectionName, kmap)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func (*postService) DeletePost(data []byte) (*mongo.DeleteResult, error) {
	temp := string(data)
	kmap := make(map[string]interface{})
	json.Unmarshal([]byte(temp), &kmap)
	result, err := mongoRepo.DeleteDocument(CollectionName, kmap)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

/*
* Private methods help in building filters thay could be utilize in specialized data retrivel

func filterFormate(data []byte) primitive.M {
	var filter primitive.M
	temp := string(data)
	k := make(map[string]string)
	log.Println("map", k)
	json.Unmarshal([]byte(temp), &k)
	for key, value := range k {
		if key == "Id" {
			objId, _ := primitive.ObjectIDFromHex(value)
			filter = bson.M{"_id": objId}
		} else {
			filter = bson.M{key: value}
		}
	}
	log.Println("filtered values=", filter)
	return filter
}


update := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 16}}}}

func updateFilter(data []byte) (primitive.M, []primitive.M) {
	var filter primitive.M
	var update []primitive.M
	i := 1
	k := make(map[string]string)
	tmp := string(data)
	json.Unmarshal([]byte(tmp), &k)
	for key, value := range k {
		if i == 1 {
			if key == "Id" {
				objId, _ := primitive.ObjectIDFromHex(value)
				filter = bson.M{"_id": objId}
			} else {
				filter = bson.M{key: value}
			}
			i++
			continue
		}
		z := bson.M{"$set": bson.M{key: value}}
		update = append(update, z)
	}
	log.Println("filter-", filter, "update=", update)
	return filter, update
}
*/
