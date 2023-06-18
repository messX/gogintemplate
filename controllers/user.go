package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/services"
)

type UserController struct {
	UserService *services.UserService
}

func (ctrl *UserController) Register(ctx *gin.Context) {
	var input dto.RegisterDto
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code: http.StatusBadRequest, Message: "Invalid request",
			Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	logger.Debugf("Register request for user %s", input.Username)
	ctrl.UserService.SetContext(ctx)
	user, regErr := ctrl.UserService.Register(&input)
	if regErr != nil {
		logger.Errorf("Unable to create user due to err %v", regErr)
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Data:    map[string]interface{}{"data": regErr.Error()}})
		return
	}
	logger.Infof("Successfully created object")
	ctx.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) Login(ctx *gin.Context) {
	var loginDto dto.RegisterDto
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code: http.StatusBadRequest, Message: "Invalid request",
			Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	logger.Debugf("Login for user %s", loginDto.Username)
	ctrl.UserService.SetContext(ctx)
	token, loginErr := ctrl.UserService.Login(&loginDto)
	if loginErr != nil {
		logger.Errorf("Unable to login user due to err %v", loginErr)
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Data:    map[string]interface{}{"data": loginErr.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func (ctrl *UserController) GetUser(ctx *gin.Context) {
	ctrl.UserService.SetContext(ctx)
	userDto, loginErr := ctrl.UserService.GetUserFromCtx()
	if loginErr != nil {
		logger.Errorf("Unable to login user due to err %v", loginErr)
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Data:    map[string]interface{}{"data": loginErr.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userDto})
}
