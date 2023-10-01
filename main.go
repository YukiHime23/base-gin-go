package main

import (
	"base-gin-go/api"
	config_caarlos0_env "base-gin-go/config/caarlos0-env"
	config_godotenv "base-gin-go/config/godotenv"
	"base-gin-go/infra"
	"base-gin-go/middleware"
	"base-gin-go/pkg/i18n"
	"context"
	"fmt"
	"log"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	fmt.Println("---- Hello world! ----")

	config_godotenv.Setup()

	cfg, err := config_caarlos0_env.New()
	if err != nil {
		log.Fatal("Load env fail")
	}

	i18n.SetupI18n()
	middleware.ValidateFunction()

	dbClient, err := infra.PostgresOpen(context.Background(), cfg)
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	server := api.SetupServer(dbClient)
	infra.Migration(dbClient)

	err = server.Router.Run(":8080")
	if err != nil {
		log.Fatal("Start fail")
	}
}
