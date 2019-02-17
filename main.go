package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rumblefrog/DBL-Webhook-Proxy/database"
	"github.com/rumblefrog/DBL-Webhook-Proxy/env"
	"github.com/rumblefrog/DBL-Webhook-Proxy/graceful"
	"github.com/rumblefrog/DBL-Webhook-Proxy/listener"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.RegisterExitHandler(graceful.GracefulExit)
}

func main() {
	var (
		dbHost       = env.Get("DB_HOST", "127.0.0.1")
		dbPort       = env.Get("DB_PORT", "3376")
		dbUsername   = env.Get("DB_USERNAME", "")
		dbPassword   = env.Get("DB_PASSWORD", "")
		databaseName = env.Get("DATABASE", "")
	)

	if dbUsername == "" || databaseName == "" {
		logrus.Fatal("Missing env variables")
	}

	database.Connect(dbHost, dbPort, dbUsername, dbPassword, databaseName)

	go listener.Start()

	logrus.Info("Listener started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Fatal("Received exit signal. Terminating.")
}
