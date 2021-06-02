package main

import (
	"database/sql"
	"github.com/go-openapi/loads"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"user_server/gen/restapi"
	"user_server/gen/restapi/operations"
	"user_server/internal/userstorage"
)

type Environment struct {
	Port         int    `split_words:"true"`
	DatabasePath string `split_words:"true"`
}

func main() {
	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("sqlite3", env.DatabasePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewUserStorageAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = env.Port

	app := userstorage.NewApp(db)
	app.ConfigureAPI(api)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
