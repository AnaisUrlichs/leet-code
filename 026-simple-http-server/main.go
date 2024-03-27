package main

//https://dev.to/stungnet/making-a-basic-http-server-with-golang-37lk

import (
	"fmt"
	"log"
	"net/http"
)

const portNum string = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func main() {
	log.Println("Starting server")

	http.HandleFunc("/", Home)

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	err := http.ListenAndServe(portNum, nil)

	if err != nil {
		log.Fatal(err)
	}
}
