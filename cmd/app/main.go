package main

import (
	"log"
	"testMod/cmd/server"
	"testMod/config"
	"testMod/database"
	"testMod/handler"
	"testMod/models"
	"testMod/pkg/rabbitMq"
	"testMod/repository"
	"testMod/service"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)

	}
	conn, err := database.Connect(config.AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(conn)

	r, err := rabbitMq.Connect(config.AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	_, err = r.DeclareQueue("otp")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
	userRepo := repository.NewUserRepository(conn)
	userservice := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userservice)

	authRepo := repository.NewAuthRepository(conn)
	authservice := service.NewAuthService(authRepo, config.AppConfig, r)
	authHandler := handler.NewAuthHandler(authservice)
	err = server.StartServer(config.AppConfig, userHandler, authHandler)
	if err != nil {
		log.Fatal(err)
	}
}
