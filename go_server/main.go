package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwilla/trickl/views"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Home().Render(r.Context(), w)
	})

	fmt.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
