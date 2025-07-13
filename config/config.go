package config

import (
	"os"
	"github.com/joho/godotenv"
)
	

type Config struct{
Database database
AppName string
AppPort string
}

type database struct{
	PostgresUser string
	PostgresHost string
	PostgresPassword string
	PostgresDb string
	PostgresPort string
	SecretKey string
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
AppConfig=cfg
return nil
}