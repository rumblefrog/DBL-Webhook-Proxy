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

	// TODO: Switch to post
	if req.Method == "GET" && len > 4 {
		authorization := req.Header.Get("Authorization")

		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing authorization header"))

			return
		}

		endpoint, err := endpoint.ParsePath(req.URL.Path)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error parsing endpoint"))
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Resource not found"))
}
