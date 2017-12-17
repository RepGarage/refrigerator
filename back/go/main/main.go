package main

import (
	"log"
	"os"

	"github.com/centrypoint/refrigerator/back/go/main/server"
)

var (
	// Info logger
	Info             *log.Logger
	api              server.API
	databaseInstance server.Database
	dbName           = os.Getenv("DATABASE_NAME")
	serverPort       = os.Getenv("SERVER_PORT")
)

func main() {
	SetupLogger("dev")
	database := databaseInstance
	e := database.Connect(dbName)
	if e != nil {
		panic(e)
	}
	server := server.Server{Port: serverPort, DB: &database}
	server.Serve(api)
}

// SetupLogger set logger prefix
func SetupLogger(env string) {
	var prefix string
	switch env {
	case "test":
		prefix = "Test: "
		break
	case "dev":
		prefix = "Info: "
		break
	}
	Info = log.New(os.Stdout, prefix, log.Ltime)
}
