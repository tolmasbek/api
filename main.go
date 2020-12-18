package main

import (
	rest_api "github.com/tolmasbek/api.git"
	"github.com/tolmasbek/api.git/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)
	server := new(rest_api.Server)
	err := server.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
