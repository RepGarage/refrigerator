package main

import (
	"os"
)

var serverPort = os.Getenv("SERVER_PORT")

func main() {
	database := Db{}
	e := database.Connect()
	if e != nil {
		panic(e)
	}
	server := Server{port: serverPort, db: &database}
	server.Serve()
}
