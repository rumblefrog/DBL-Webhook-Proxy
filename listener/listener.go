package listener

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var Server *http.Server

func Start() {
	Server = &http.Server{
		Addr:    ":8080",
		Handler: &VoteHandler{},
	}

	if err := Server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithField("error", err).Fatal("Unable to serve http server")
	}
}
