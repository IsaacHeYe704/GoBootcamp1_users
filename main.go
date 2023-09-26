package main

import (
	"bootcam1_users/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/users", handlers.PostUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.PutUsers).Methods("PUT")

	fmt.Println("LISTENING TO PORT 3000")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))

}
