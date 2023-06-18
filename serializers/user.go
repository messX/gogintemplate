package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/models"
)

const _dateFormat string = "2006-01-02T15:04:05.999Z"

type UserSerializer struct {
	Ctx  *gin.Context
	User models.User
}

func (obj *UserSerializer) Response() *dto.UserDto {
	dataDto := dto.UserDto{
		ID:        obj.User.Id,
		Username:  obj.User.Username,
		CreatedAt: obj.User.CreatedAt.Format(_dateFormat),
		UpdatedAt: obj.User.UpdatedAt.Format(_dateFormat),
	}
	return &dataDto
}

func (obj *UserSerializer) SetContext(c *gin.Context) {
	obj.Ctx = c
}

func NewUserSerializer() *UserSerializer {
	return &UserSerializer{}
}
