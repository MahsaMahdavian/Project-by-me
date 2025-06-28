package database

import (
	"testMod/config"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)


func Connect(config config.Config )(*sql.DB, error){


conn, err:=sql.Open("postgres",fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable",
config.Database.PostgresUser, 
config.Database.PostgresDb,
config.Database.PostgresPassword,
config.Database.PostgresPort,
))

return conn,err
}

func Close( conn *sql.DB) error{
return conn. Close ()
}

func Ping (conn *sql.DB)error{
return conn.Ping ()
}

func ExampleQuery (conn *sql.DB) (error, string){
var res string
err := conn. QueryRow("SELECT version()").Scan (&res)
return err, res
}