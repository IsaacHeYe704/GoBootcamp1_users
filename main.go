package main

import (
	"bootcam1_users/handlers"
	"bootcam1_users/structures"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func testRedis() {
	port := goDotEnvVariable("REDIS_PORT")
	connection := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	structures.NewRedisStorage(connection)
}
func main() {
	router := mux.NewRouter()
	//leer que storage usar segun .env

	testRedis()

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

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
