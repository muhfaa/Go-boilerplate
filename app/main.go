package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/muhfaa/Go-boilerplate/config"
	"github.com/muhfaa/Go-boilerplate/modules/mysql"
)

func newDatabaseConnection() *sqlx.DB {
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetConfig().Mysql.User,
		config.GetConfig().Mysql.Password,
		config.GetConfig().Mysql.Host,
		config.GetConfig().Mysql.Port,
		config.GetConfig().Mysql.Name)

	db := mysql.NewDatabaseConnection(uri)

	return db

}

func main() {
	// initialize database connection based on given config
	// db := newDatabaseConnection()

	e := echo.New()
	e.HideBanner = true

	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().BackendPort)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
