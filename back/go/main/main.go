package main

import (
	"log"
	"os"
)

var (
	// Info logger
	Info             *log.Logger
	api              API
	databaseInstance Database
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
	server := Server{port: serverPort, db: &database}
	server.Serve(api, &databaseInstance)
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
