package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		Hello("World").Render(req.Context(), w)
	})
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
