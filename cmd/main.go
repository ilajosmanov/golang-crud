package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	s := new(todo.Server)

	if err := s.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}