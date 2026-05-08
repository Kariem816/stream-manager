package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kariem816/stream-manager/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Http)
	http.HandleFunc("/{streamId}", handlers.Stream)
	fmt.Println("Stream Manager running on http://localhost:13118")
	log.Fatal(http.ListenAndServe(":13118", nil))
}
