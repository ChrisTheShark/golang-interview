package handlers

import "net/http"

type SimpleStruct struct {
	Reference   string `json:"reference"`
	HiddenValue string `json:"hidden"`
}

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello, World!"))
}
