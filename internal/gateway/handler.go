package gateway

import (
	"log"
	"net/http"
)

// UserHandler handles /users REST endpoint and forwards to User gRPC service
func UserHandler(cfg Config, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Connect to UserService gRPC and forward request
		logger.Println("Received request at /users (placeholder)")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Not implemented"}`))
	}
}
