package main

import (
	"fmt"
	"log"
	"testMod/cmd/server"
	"testMod/config"
	"testMod/database"
	"testMod/handler"
	"testMod/repository"
	"testMod/service"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)

	}
	conn, err := database.Connect(config.AppConfig)
	fmt.Println(conn.Ping())
	defer database.Close(conn)

	userRepo := repository.NewUserRepository(conn)
	userservice := service.NewUserService(userRepo)

	userHandler:=handler.NewUserHandler(userservice)

	err = server.StartServer(config.AppConfig, userHandler)
	if err != nil {
		log.Fatal(err)
	}
}
