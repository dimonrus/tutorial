package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Init register system sub routes
func Init(apiRoute *mux.Router, middleware ...mux.MiddlewareFunc) {
	// Database activities
	apiRoute.HandleFunc("/activity", CurrentActivityHandler).Methods(http.MethodGet)
	// Use middleware
	apiRoute.Use(middleware...)
}
