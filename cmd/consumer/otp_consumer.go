package main

import (
	"fmt"
	"log"
	// "testMod/cmd/server"
	"testMod/config"
	"testMod/database"
	"testMod/pkg/rabbitMq"
	"testMod/repository"
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

	userRepo := repository.NewUserRepository(conn)

	messages, err := r.ConsumeMessages("otp")
	if err != nil {
		log.Fatal(err)
	}
	for msg := range messages {
		user,err:=userRepo.Find(string(msg.Body))
		if err != nil {
			log.Println("Error finding user:", err)
			continue
		}
		//send sms
		fmt.Println(user.Mobile)
		msg.Ack(true)
	}
}

