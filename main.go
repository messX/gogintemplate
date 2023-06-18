package main

import (
	"time"

	"github.com/messx/gogintemplate/config"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/messx/gogintemplate/routers"
	"github.com/spf13/viper"
)

// @title           Swagger Example API
// @version         1.0

// @contact.name   API Support

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	err := config.SetupConfig()
	if err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	//run database
	config.ConnectDB()
	router := routers.Routes()
	// initialise SQS watchers
	go config.InitAndProcessSQS()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
