package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/errors"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/models"
	"github.com/messx/gogintemplate/repository"
	"github.com/messx/gogintemplate/serializers"
)

var userRepository repository.UserRepository = *repository.NewUserRepository()

type UserService struct {
	c *gin.Context
}

func _checkUserNameAvail(ctx *gin.Context, username string) (*models.User, error) {
	return userRepository.FindWithUserName(ctx, username)
}

func (obj *UserService) Register(register *dto.RegisterDto) (*dto.UserDto, error) {
	logger.Debugf("Check if user %s already exists", register.Username)
	user, err := _checkUserNameAvail(obj.c, register.Username)
	if user != nil {
		logger.Debugf("Error %v", err)
		logger.Errorf("Username %s already taken", register.Username)
		userErr := &errors.UserTakenError{
			Username: register.Username,
		}
		return nil, userErr
	}
	u := models.User{
		Username:  register.Username,
		Password:  register.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, insErr := userRepository.InsertOne(obj.c, &u)
	if insErr != nil {
		logger.Fatalf("Unable to insert model %v", insErr)
		return nil, insErr
	} else {
		dataSerializer := serializers.UserSerializer{Ctx: obj.c, User: u}
		return dataSerializer.Response(), nil
	}
}

func (obj *UserService) Login(loginDto *dto.RegisterDto) (string, error) {
	user, err := userRepository.FindWithUserName(obj.c, loginDto.Username)
	if err != nil {
		logger.Debugf("Error %v", err)
		logger.Errorf("Username invalid", loginDto.Username)
		userErr := &errors.UserNameInvalid{
			Username: loginDto.Username,
		}
		return "", userErr
	}
	pwdErr := userRepository.VerifyPassword(loginDto.Password, user)
	if pwdErr != nil {
		logger.Debugf("Error %v", err)
		logger.Errorf("Password invalid", loginDto.Username)
		userErr := &errors.UserPasswordInvalid{
			Username: loginDto.Username,
		}
		return "", userErr
	}
	return helpers.GenerateToken(user.Id.Hex())
}

func (obj *UserService) GetUserFromCtx() (*dto.UserDto, error) {
	userId, err := helpers.ExtractTokenID(obj.c)
	if err != nil {
		logger.Debugf("Error %v", err)
		return nil, err
	}
	user, userErr := userRepository.GetByUserId(obj.c, userId)
	if userErr != nil {
		logger.Debugf("Error %v", userErr)
		return nil, userErr
	}
	dataSerializer := serializers.UserSerializer{Ctx: obj.c, User: *user}
	return dataSerializer.Response(), nil
}

func NewUserService() *UserService {
	return &UserService{}
}

func (obj *UserService) SetContext(c *gin.Context) {
	obj.c = c
}
