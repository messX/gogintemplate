package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/messx/gogintemplate/infra/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client instance
var DB *mongo.Client = ConnectDB()

func ConnectDB(env ...string) *mongo.Client {
	logger.Debugf("Connecting to mongo db %s", viper.GetString("MONGOURI"))
	var mongouri string = viper.GetString("MONGOURI")
	if len(mongouri) == 0 {
		logger.Debugf("Setting up env incase of not load")
		SetupConfig(env...)
		mongouri = viper.GetString("MONGOURI")
	}
	// fmt.Printf(viper.GetString("MONGOURI"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongouri))
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://testuser:password@localhost:27017/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database(viper.GetString("MONGO_DB_NAME")).Collection(collectionName)
	return collection
}

// creating database collections
func CreateCollection(client *mongo.Client, collectionName string) error {
	ctx := context.Background()
	err := client.Database(viper.GetString("MONGO_DB_NAME")).CreateCollection(ctx, collectionName)
	return err
}

func DropCollection(client *mongo.Client, collectionName string) error {
	ctx := context.Background()
	err := client.Database(viper.GetString("MONGO_DB_NAME")).Collection(collectionName).Drop(ctx)
	return err
}

func ListCollectionNames(client *mongo.Client) ([]string, error) {
	ctx := context.Background()
	list, err := client.Database(viper.GetString("MONGO_DB_NAME")).ListCollectionNames(ctx, bson.M{})
	return list, err
}
