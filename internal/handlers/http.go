package handlers

import "net/http"

func Http(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Hello World!\"}"))
}
