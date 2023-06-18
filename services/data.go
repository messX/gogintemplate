package services

import (
	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/models"
	"github.com/messx/gogintemplate/repository"
	"github.com/messx/gogintemplate/serializers"
)

var dataRepository repository.DataRepository = *repository.NewDataRepository()

type DataService struct {
	c *gin.Context
}

func (obj *DataService) GetAll() ([]*dto.DataDto, error) {
	results, err := dataRepository.FindAll(obj.c)
	if err != nil {
		logger.Fatalf("Unable to fecth data %v", err)
		return nil, err
	} else {
		var response []*dto.DataDto
		for _, res := range results {
			dataSerializer := serializers.DataSerializer{Ctx: obj.c, Data: res}
			response = append(response, dataSerializer.Response())
		}
		return response, nil
	}

}

func (obj *DataService) Create(data *models.Data) (*dto.DataDto, error) {
	_, err := dataRepository.InsertOne(obj.c, data)
	if err != nil {
		logger.Fatalf("Unable to insert model %v", err)
		return nil, err
	} else {
		dataSerializer := serializers.DataSerializer{Ctx: obj.c, Data: *data}
		return dataSerializer.Response(), nil
	}
}

func NewDataService() *DataService {
	return &DataService{}
}

func (obj *DataService) SetContext(c *gin.Context) {
	obj.c = c
}
