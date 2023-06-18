package tests

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/messx/gogintemplate/models"
)

func TestAddEntry(t *testing.T) {
	newData := models.Data{
		Name: "Test",
	}
	writer := makeRequest("POST", "/api/v1/data", newData)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAllEntris(t *testing.T) {
	writer := makeRequest("GET", "api/v1/data", nil)
	assert.Equal(t, http.StatusOK, writer.Code)
}
