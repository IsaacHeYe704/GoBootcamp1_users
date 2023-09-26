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
	//leer que storage usar segun .env

	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/users", handlers.PostUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.PutUsers).Methods("PUT")

	fmt.Println("LISTENING TO PORT 3000")
	// Bind to a port and pass our router in
	//Errores muy importantes --> generar panic que detiene la app por completo
	log.Fatal(http.ListenAndServe(":3000", router))

}
