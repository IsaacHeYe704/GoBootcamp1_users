package main

import (
	"bootcam1_users/db"
	"bootcam1_users/handlers"
	"bootcam1_users/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	router := mux.NewRouter()
	var storage db.Storage
	//Read .env to choose if we should use localstorage or redis
	switch goDotEnvVariable("STORAGE") {
	case "Redis":
		fmt.Println("USING REDIS STORAGE...")
		var connection = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf(goDotEnvVariable("REDIS_ADDR")),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		storage = db.NewRedisStorage(connection)
	default:
		fmt.Println("USING LOCAL STORAGE...")
		storage = db.NewLocalStorage()
	}
	//declare the user service injecting the storage dependency
	var usersService = service.NewUserService(storage)

	//ROUTER
	router.HandleFunc("/users", handlers.GetAllUsers(usersService)).Methods("GET")
	router.HandleFunc("/users", handlers.PostUser(usersService)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserById(usersService)).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteUsers(usersService)).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.PutUsers(usersService)).Methods("PUT")

	fmt.Println("LISTENING TO PORT 3000")
	// // Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Auxiliar function to get a .env var
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
