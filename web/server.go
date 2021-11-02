package web

import (
	"log"
	"net/http"
)

func Server() {
	get_port()
	
	// Route: /
	http.HandleFunc("/", root_handler)
	log.Fatal(http.ListenAndServe(":20000", nil))
}