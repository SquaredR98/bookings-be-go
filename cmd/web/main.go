package main

import (
	"fmt"
	"net/http"

	"github.com/SquaredR98/bookings-be/pkg/handlers"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server running on PORT: %s", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
