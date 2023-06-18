package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/dto"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/models"
	"github.com/messx/gogintemplate/services"
)

type MainController struct {
	DataService *services.DataService
}

// HeaderExample godoc
//
//	@Summary		get request example
//	@Description	get all data
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Success		200				{string}	string	"answer"
//	@Failure		400				{string}	string	"ok"
//	@Failure		404				{string}	string	"ok"
//	@Failure		500				{string}	string	"ok"
//	@Router			/api/v1/data [get]
func (ctrl *MainController) GetAllData(ctx *gin.Context) {

	var data []*dto.DataDto
	ctrl.DataService.SetContext(ctx)
	// defer cancel()
	results, err := ctrl.DataService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Data:    map[string]interface{}{"data": err.Error()}})
		return
	}
	data = results
	ctx.JSON(http.StatusOK,
		helpers.Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data},
	)
}

// PostExample godoc
//
//	@Summary		post request example
//	@Description	post request example
//	@Accept			json
//	@Produce		plain
//	@Param			message	body		models.Data	true	"Sample data"
//	@Success		200		{string}	string			"success"
//	@Failure		500		{string}	string			"fail"
//	@Router			/api/v1/data [post]
func (ctrl *MainController) Create(ctx *gin.Context) {
	logger.Debugf("Recieved request to create data")
	example := new(models.Data)
	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code: http.StatusBadRequest, Message: "Invalid request",
			Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctrl.DataService.SetContext(ctx)
	res, err := ctrl.DataService.Create(example)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Data:    map[string]interface{}{"data": err.Error()}})
		return
	}
	logger.Infof("Successfully created object")
	ctx.JSON(http.StatusCreated, res)

}
