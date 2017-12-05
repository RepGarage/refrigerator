package main

import (
	"os"
)

var serverPort = os.Getenv("SERVER_PORT")

// DatabaseInterfaceInstance DatabaseInterface instance
var DatabaseInterfaceInstance DatabaseInterface

func main() {
	database := DatabaseInterfaceInstance
	e := database.Connect()
	if e != nil {
		panic(e)
	}
	server := Server{port: serverPort, db: &database}
	server.Serve()
}
