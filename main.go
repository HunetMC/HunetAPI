package main

import (
	"fmt"
	"log"
	"net/http"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	resp := "200 OK"
	fmt.Fprint(w, resp)
}

func main() {
	// Route: /
	http.HandleFunc("/", root_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}