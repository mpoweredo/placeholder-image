package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"placeholder-image/config"
	"placeholder-image/routes"
)

func main() {
	port, err := config.LoadPORT()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	routes.PlaceholderRoutes(r)

	http.Handle("/", r)
	fmt.Println("listening on port: ", ":"+port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
