package main

import (
	"log"
	"net/http"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ABC"))
}

func main() {
	http.HandleFunc("/", sayHi)
	port := ":3000"

	log.Print(`Running at http://localhost` + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
