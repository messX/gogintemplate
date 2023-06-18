package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/models"
)

type DataSerializer struct {
	Ctx  *gin.Context
	Data models.Data
}

func (obj *DataSerializer) Response() *dto.DataDto {
	dataDto := dto.DataDto{
		ID:   obj.Data.Id,
		Name: obj.Data.Name,
	}
	return &dataDto
}

func (obj *DataSerializer) SetContext(c *gin.Context) {
	obj.Ctx = c
}

func NewDataSerializer() *DataSerializer {
	return &DataSerializer{}
}
