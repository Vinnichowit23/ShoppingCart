package library

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getusers).Methods("GET")
	r.HandleFunc("/users/{id}", getuser).Methods("GET")
	r.HandleFunc("/users", addusers).Methods("POST")
	r.HandleFunc("/users/{id}", updateusers).Methods("PATCH")
	r.HandleFunc("/users/{id}", deleteusers).Methods("DELETE")

	fmt.Printf("starting server at 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
