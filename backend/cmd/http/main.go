package main

import (
	"net/http"

	handlers "github.com/fiyintheslim/chat-app-v2.git/backend/internal/HTTPHandlers"
	db_conn "github.com/fiyintheslim/chat-app-v2.git/backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db_conn.ConnectDB()
	http.ListenAndServe(":8080", handlers.HandleHTTPRoutes())
}