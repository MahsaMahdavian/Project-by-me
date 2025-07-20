package config

import (
	"os"
	"github.com/joho/godotenv"
)
	

type Config struct{
Database database
AppName string
AppPort string
Rabbitmq rabbitmq
}

type database struct{
	PostgresUser string
	PostgresHost string
	PostgresPassword string
	PostgresDb string
	PostgresPort string
	SecretKey string
}
type rabbitmq struct{
	RabbitmqUserName string
	RabbitmqPassword string
	RabbitmqHost string
	RabbitmqPort string
}

var AppConfig Config

func LoadConfig()error{
var cfg Config


err:=godotenv.Load(".env")
if err!=nil{
	return err
}
cfg.AppName=os.Getenv("App_Name")
cfg.AppPort=os.Getenv("App_Port")
cfg.Database.PostgresUser=os.Getenv("POSTGRES_USER")
cfg.Database.PostgresPassword=os.Getenv("POSTGRES_PASSWORD")
cfg.Database.PostgresHost=os.Getenv("POSTGRES_HOST")
cfg.Database.PostgresPort=os.Getenv("POSTGRES_PORT")
cfg.Database.PostgresDb=os.Getenv("POSTGRES_DB")
cfg.Database.SecretKey=os.Getenv("SECRET_KEY")
cfg.Rabbitmq.RabbitmqUserName=os.Getenv("RABBITMQ_USERNAME")
cfg.Rabbitmq.RabbitmqPassword=os.Getenv("RABBITMQ_PASSWORD")
cfg.Rabbitmq.RabbitmqHost=os.Getenv("RABBITMQ_HOST")
cfg.Rabbitmq.RabbitmqPort=os.Getenv("RABBITMQ_PORT")
AppConfig=cfg
return nil
}