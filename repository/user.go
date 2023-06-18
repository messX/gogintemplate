package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/config"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	userRepository := new(UserRepository)
	userRepository.collection = config.GetCollection(config.DB, "users")
	return userRepository
}

func (_m *UserRepository) InsertOne(ctx *gin.Context, data *models.User) (*mongo.InsertOneResult, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	data.Password = string(hashedPassword)
	return _m.collection.InsertOne(ctx, data)
}

func (_m *UserRepository) FindWithUserName(ctx *gin.Context, username string) (*models.User, error) {
	filter := bson.M{"username": username}
	var result models.User
	err := _m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		logger.Errorf("Unable to fetch user %v", err)
		return nil, err
	}
	return &result, err
}

func (_m *UserRepository) GetByUserId(ctx *gin.Context, userId string) (*models.User, error) {
	objId, parErr := primitive.ObjectIDFromHex(userId)
	if parErr != nil {
		log.Println("Invalid id")
	}
	filter := bson.M{"id": objId}
	var result models.User
	err := _m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		logger.Errorf("Unable to fetch user %v", err)
		return nil, err
	}
	return &result, err
}

func (_m *UserRepository) VerifyPassword(password string, data *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
}
