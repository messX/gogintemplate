package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/config"
	"github.com/messx/gogintemplate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataRepository struct {
	collection *mongo.Collection
}

func NewDataRepository() *DataRepository {
	dataRepository := new(DataRepository)
	dataRepository.collection = config.GetCollection(config.DB, "data")
	return dataRepository
}

func (_m *DataRepository) FindAll(ctx *gin.Context) ([]models.Data, error) {
	var data []models.Data
	logger := log.Default()
	results, err := _m.collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Fatal("UNable to fetch data")
		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleData models.Data
		int_err := results.Decode(&singleData)
		if int_err != nil {
			logger.Fatal("UNable to fetch data")
			return nil, int_err
		}
		data = append(data, singleData)
	}
	return data, nil
}

func (_m *DataRepository) InsertOne(ctx *gin.Context, data *models.Data) (*mongo.InsertOneResult, error) {
	return _m.collection.InsertOne(ctx, data)
}
