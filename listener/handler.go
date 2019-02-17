package listener

import (
	"net/http"

	"github.com/rumblefrog/DBL-Webhook-Proxy/endpoint"
)

type VoteHandler struct{}

func (v *VoteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		w.Write([]byte("DBL Webhook Proxy"))

		return
	}

	len := len(req.URL.Path)

	if req.Method == "GET" && len > 4 {
		endpoint, err := endpoint.ParsePath(req.URL.Path)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("Error parsing endpoint"))
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Resource not found"))
}
