package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Setup() {

	mongoUrl := "mongodb://127.0.0.1:27017"
	MongoDbName = "mytest"

	var err error
	clientOptions := options.Client().ApplyURI(mongoUrl).SetConnectTimeout(5 * time.Second).SetMaxPoolSize(32)
	//clientOptions.Monitor = otelmongo.NewMonitor()

	// 连接到MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//CheckCollections()
}

// func CheckCollections() {
// 	GetMongoClient()

// 	collectionNames, err := client.Database(MongoDbName).ListCollectionNames(context.TODO(), bson.M{})
// 	if err == nil {
// 		collectionNameMap := make(map[string]bool)
// 		for _, collectionName := range collectionNames {
// 			collectionNameMap[collectionName] = true
// 		}
// 		if !collectionNameMap[PostTableName] {
// 			//create the post collection
// 			createPostCollectionIndexes()
// 		}

// 		if !collectionNameMap[CommentTableName] {
// 			createCommentCollectionIndexes()
// 		}
// 	}
// }

func GetMongoClient() *mongo.Client {
	if client == nil {
		Setup()
	}
	return client
}
