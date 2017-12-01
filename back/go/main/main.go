package main

func main() {
	database := Db{}
	e := database.Connect()
	if e != nil {
		panic(e)
	}
	server := Server{port: ":8080", db: &database}
	server.Serve()
}
