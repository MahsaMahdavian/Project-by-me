package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testMod/config"
)

func Connect(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.Database.PostgresHost,
		config.Database.PostgresUser,
		config.Database.PostgresPassword,
		config.Database.PostgresDb,
		config.Database.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return db, err
	}
	return db, nil
}

func Close(conn *gorm.DB) error {
	sqlDb, err := conn.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}

func Ping(conn *gorm.DB) error {
	sqlDb, err := conn.DB()
	if err != nil {
		return err
	}
	return sqlDb.Ping()
}
