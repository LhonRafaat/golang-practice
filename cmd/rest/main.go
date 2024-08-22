package main

import (
	"lhon/postgres-rest/internal/repository"
	"lhon/postgres-rest/internal/server"
)

func main() {


	repository.InitDB()
	defer repository.CloseDB()
	
	server.StartServer()

}