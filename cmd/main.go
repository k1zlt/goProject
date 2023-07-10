package main

import (
	"first/internal/handler"
	"first/internal/repository"
	"first/internal/repository/postgres"
	"first/internal/server"
	"first/internal/service"
	"first/logs"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	file := logs.SetUpLog()
	defer file.Close()

	db, err := postgres.ConnectToPostgresDB(postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12345",
		DBName:   "online_school",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("FATAL: failed to connect database: %s", err)
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handlers := handler.NewHandler(srvc)
	srv := new(server.Server)

	router := handlers.InitRoutes()
	if err := srv.Run("8080", router); err != nil {
		log.Fatalf("FATAL: error occurred while running the server: %s", err)
	}
}
