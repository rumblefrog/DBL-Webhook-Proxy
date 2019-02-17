package listener

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	Server         *http.Server
	DefaultTimeout = time.Second * 3
)

func Start() {
	Server = &http.Server{
		Addr:              ":8080",
		Handler:           &VoteHandler{},
		ReadHeaderTimeout: DefaultTimeout,
	}

	if err := Server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithField("error", err).Fatal("Unable to serve http server")
	}
}
