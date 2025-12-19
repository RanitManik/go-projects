package routes

import (
	"github.com/gorilla/mux"
	"github.com/RanitManik/go-projects/03-book-management-system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
}
