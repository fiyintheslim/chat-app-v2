package api_handler

import (
	"io"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Welcome to hello")
}
