package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/dbx"
	"red-envelope/infra"
	"time"
)

var database *dbx.Database

func DbxDataBase() *dbx.Database {
	return database
}

type DbxDataBaseStarter struct {
	infra.BaseStarter
}

func (d *DbxDataBaseStarter)Setup(ctx infra.StarterContext)  {
	settings := dbx.Settings{
	       DriverName: "mysql",
	       User:       "red",
	       Password:   "123",
	       Host:       "127.0.0.1:3306",
	       //Host:            "172.16.1.248:3306",
	       Database:        "po",
	       MaxOpenConns:    10,
	       MaxIdleConns:    2,
	       ConnMaxLifetime: time.Minute * 30,
	       LoggingEnabled:  true,
	       Options: map[string]string{
	           "charset":   "utf8",
	           "parseTime": "true",
	       },
	   }

	dbx,err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	fmt.Println("dbx ping:",dbx.Ping())
	database =dbx
}
