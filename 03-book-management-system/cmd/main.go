package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RanitManik/go-projects/03-book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	addr := ":8080"
	fmt.Println("Server running on http://localhost" + addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
