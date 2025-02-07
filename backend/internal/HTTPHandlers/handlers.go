package handlers

import (
	"net/http"

	api_handler "github.com/fiyintheslim/chat-app-v2.git/backend/internal/HTTPHandlers/api"
)

func HandleHTTPRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", api_handler.HandleHello)

	return mux
}
