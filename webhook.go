package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var payload map[string]interface{}
		json.NewDecoder(r.Body).Decode(&payload)
		response := map[string]string{"message": "Hola Mundo!", "received": fmt.Sprint(payload)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	http.ListenAndServe(":8080", nil)
}
