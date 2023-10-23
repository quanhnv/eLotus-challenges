package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/quanhnv/eLotus-challenges/middlewares"
	"github.com/quanhnv/eLotus-challenges/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	//Hello world
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	//For auth
	http.HandleFunc("/register", routes.Register)
	http.HandleFunc("/login", routes.Login)

	//For upload file
	http.HandleFunc("/upload", middlewares.CheckJwt(routes.Upload))

	port := os.Getenv("APP_PORT")
	//Server port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
