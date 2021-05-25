package main

import (
	"log"

	"ofortuna/app"
)

//  network ip/port bind info
const address string = ":1999"

func main() {
	log.Print("Starting O Fortuna!")
	log.Fatal(app.ListenAndServe(address))
}
