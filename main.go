package main

import (
	"fmt"
	"log"
	"net/http"
)

// with thanks to Carl Orff
// https://en.wikipedia.org/wiki/O_Fortuna
const ofortuna string = `O Fortuna!

velut luna
statu variabilis,
semper crescis
aut decrescis;
vita detestabilis
nunc obdurat
et tunc curat
ludo mentis aciem,
egestatem,
potestatem
dissolvit ut glaciem.
`

//  network ip/port bind info
const server string = ":1999"

func main() {
	log.Print("Starting O Fortuna!")

	http.HandleFunc("/", Handler)
	http.ListenAndServe(server, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetFortune())
}
func GetFortune() string {
	return (ofortuna)
}
