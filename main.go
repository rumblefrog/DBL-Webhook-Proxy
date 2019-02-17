package main

import (
	"github.com/rumblefrog/DBL-Webhook-Proxy/database"
	"github.com/rumblefrog/DBL-Webhook-Proxy/env"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		dbHost       = env.Get("DB_HOST", "127.0.0.1")
		dbPort       = env.Get("DB_PORT", "3376")
		dbUsername   = env.Get("DB_USERNAME", "")
		dbPassword   = env.Get("DB_PASSWORD", "")
		databaseName = env.Get("DATABASE", "")
	)

	if dbUsername == "" || dbPassword == "" || databaseName == "" {
		logrus.Fatal("Missing env variables")
	}

	database.Connect(dbHost, dbPort, dbUsername, dbPassword, databaseName)
}
