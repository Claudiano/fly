package main

import (
	"fly/routers"
	"fly/utils"
	"fmt"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
func main() {

	utils.ValidarServicos()

	fmt.Println(time.Now())

	fmt.Println("Iniciando servidor")
	routers.InitServer()
}
