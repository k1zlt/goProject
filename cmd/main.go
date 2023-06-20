package main

import (
	"first/internal/handler"
	"first/internal/repository"
	"first/internal/server"
	"first/internal/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := repository.ConnectToPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12345",
		DBName:   "online_school",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handlers := handler.NewHandler(srvc)
	srv := new(server.Server)

	router := handlers.InitRoutes()
	if err := srv.Run("8080", router); err != nil {
		log.Fatal("error occurred while running the server")
	}
	log.Println("Server started")
}
