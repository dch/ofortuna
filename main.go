package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dch/ofortuna/fortunes"
)

//  network ip/port bind info
const server string = ":1999"

func main() {
	log.Print("Starting O Fortuna!")

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(server, nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fortunes.GetRandomFortune())
}
