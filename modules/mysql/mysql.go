package mysql

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type DatabaseConnection struct {
	MySQL *sqlx.DB
}

func NewDatabaseConnection(uri string) *sqlx.DB {

	mysqlDB, err := sqlx.Open("mysql", uri)
	if err != nil {
		log.Info(err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = mysqlDB.PingContext(ctx); err != nil {
		log.Info(err)
		panic(err)
	}

	return mysqlDB

}
