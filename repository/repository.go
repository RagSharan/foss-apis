package repository

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoMongo struct{}

type IMongoRepository interface {
	FindOne(collectionName string, kmap map[string]interface{}) (bson.M, error)
	FindAll(collectionName string, kmap map[string]interface{}) ([]bson.M, error)
	FindList(collectionName string, kmap map[string]interface{}) ([]bson.M, error)
	Create(callectionName string, document interface{}) (*mongo.InsertOneResult, error)
	CreateMany(collectionName string, document []interface{}) (*mongo.InsertManyResult, error)
	UpdateById(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error)
	UpdateOne(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error)
	UpdateMany(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error)
	DeleteDocument(collectionName string, kmap map[string]interface{}) (*mongo.DeleteResult, error)
}

func ObjIMongoRepository() IMongoRepository {
	return &repoMongo{}
}

func (*repoMongo) FindOne(collectionName string, kmap map[string]interface{}) (bson.M, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter := formateFilter(kmap)
	log.Println("repoMongo:", filter)
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	log.Println("result value", result)
	return result, err
}

/**
* Never Use this function in production application
* curser.All will load whole collection in memory
* suppose we have millions record then this function is going to fetch all records
**/
func (*repoMongo) FindAll(collectionName string, kmap map[string]interface{}) ([]bson.M, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter := formateFilter(kmap)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Println(err)
	}
	return result, err
}

/**
* This function will provide List of objects in form of cursor which needs to decode
* If filter is null here it will fetch whole database
**/
func (*repoMongo) FindList(collectionName string, kmap map[string]interface{}) ([]bson.M, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	findOptions := options.Find()
	filter := formateFilter(kmap)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Println(err)
	}
	var result []bson.M
	for cursor.Next(context.TODO()) {
		var profile bson.M
		cursor.Decode(&profile)
		result = append(result, profile)
	}
	return result, err
}

func (*repoMongo) Create(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func (*repoMongo) CreateMany(collectionName string, data []interface{}) (*mongo.InsertManyResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	result, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

/* In this method update object will hve specific settings*/
func (*repoMongo) UpdateById(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) UpdateOne(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) UpdateMany(collectionName string, kmap map[string]interface{}) (*mongo.UpdateResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter, update := formateUpdate(kmap)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (*repoMongo) DeleteDocument(collectionName string, kmap map[string]interface{}) (*mongo.DeleteResult, error) {
	database, client := connectDB()
	defer closeConnection(client)
	collection := database.Collection(collectionName)
	filter := formateFilter(kmap)
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
	return deleteResult, err
}

/**
* This function will provide connection with the DB
**/
func connectDB() (*mongo.Database, *mongo.Client) {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	database := client.Database("FooApi")
	fmt.Println("Connected to MongoDB!")
	return database, client
}
func closeConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connection to MongoDB instance is closed.")
}

func formateUpdate(kmap map[string]interface{}) (primitive.M, []primitive.M) {
	var filter primitive.M
	var update []primitive.M
	i := 0
	for key, value := range kmap {
		if i == 0 {
			zmap := map[string]interface{}{key: value}
			filter = formateFilter(zmap)
			i = 1
			continue
		}
		zmap := map[string]interface{}{key: value}
		var temp primitive.M
		keyRecursion(&temp, &zmap)
		final := bson.M{"$set": temp} //$set could be replaced by other methods
		update = append(update, final)
	}
	return filter, update
}

func formateFilter(kmap map[string]interface{}) primitive.M {
	var filter primitive.M
	for key, value := range kmap {
		if key == "_id" {
			str := fmt.Sprintf("%v", value)
			objId, _ := primitive.ObjectIDFromHex(str)
			filter = bson.M{"_id": objId}
		} else {
			keyRecursion(&filter, &kmap)
		}
	}
	return filter
}

func keyRecursion(filter *primitive.M, kmap *map[string]interface{}, tempKey ...string) {
	for key, value := range *kmap {
		if reflect.TypeOf(value).Kind() != reflect.Map {
			if tempKey != nil {
				key = tempKey[0] + "." + key
			}
			*filter = bson.M{key: value}
			fmt.Println(*filter)
		} else {
			tempKey := key
			tempMap := value.(map[string]interface{})
			keyRecursion(filter, &tempMap, tempKey)
		}
	}
}
