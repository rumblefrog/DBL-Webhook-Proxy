package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var Connection *sql.DB

func Connect(host, port, username, password, database string) {
	c := mysql.NewConfig()

	c.User = username
	c.Passwd = password
	c.DBName = database

	c.Collation = "utf8mb4_general_ci"

	c.InterpolateParams = true

	c.ParseTime = true

	c.Addr = host + ":" + port

	var err error

	Connection, err = sql.Open("mysql", c.FormatDSN())

	if err != nil {
		logrus.Fatal("Unable to open database connection")
	}
}
