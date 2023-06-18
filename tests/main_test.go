package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/config"
	"github.com/messx/gogintemplate/controllers"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/services"
)

var CollectionList = [...]string{"data"}

const configFile string = "../envs/test.env"

func TestMain(m *testing.M) {
	log.Printf("Inside main")
	logger.Errorf("Inside main")
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)

}

func router() *gin.Engine {
	router := gin.Default()
	ctrl := controllers.MainController{
		DataService: services.NewDataService(),
	}
	publicRoutes := router.Group("api/v1")
	publicRoutes.POST("/data", ctrl.Create)

	return router
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {

	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	// if isAuthenticatedRequest {
	// 	request.Header.Add("Authorization", "Bearer "+bearerToken())
	// }
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func teardown() {
	log.Printf("Testing teardown")
	tearDowmDatabase()
}

func tearDowmDatabase() {
	db := config.DB
	for _, collection := range CollectionList {
		err := config.DropCollection(db, collection)
		if err != nil {
			log.Default().Fatalf("Error in dropping collection %v", err)
		}
	}
}

func setup() {
	log.Printf("Testing setup")
	setUpTestConfig()
	setUpDatabase()
}

func setUpDatabase() {
	db := config.DB
	for _, collection := range CollectionList {
		colList, _ := config.ListCollectionNames(db)
		for _, col := range colList {
			config.DropCollection(db, col)
		}
		err := config.CreateCollection(db, collection)
		if err != nil {
			log.Default().Fatalf("Error in creating collection %v", err)
		}
	}
}

func setUpTestConfig() {

	config.SetupConfig(configFile)
}
