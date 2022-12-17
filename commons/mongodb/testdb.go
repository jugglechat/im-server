package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	ID    primitive.ObjectID `json:"-" bson:"_id"`
	Title string             `json:"title" bson:"title"`
}

type TestAdd struct {
	Title string `json:"title" bson:"title"`
}

const TestTableName = "tests"

func InsertOneTest(post TestAdd) (string, error) {
	collection := getTestCollection()
	if collection != nil {
		rst, err := collection.InsertOne(context.TODO(), post)
		if err != nil {
			return "", err
		} else {
			return rst.InsertedID.(primitive.ObjectID).Hex(), nil
		}
	} else {
		return "", &MongoError{
			Msg: "Failed insert into " + TestTableName,
		}
	}
}
func getTestCollection() *mongo.Collection {
	mongoClinet := GetMongoClient()

	if mongoClinet != nil {
		return mongoClinet.Database(MongoDbName).Collection(TestTableName)
	}
	return nil
}
