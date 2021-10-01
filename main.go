package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"net/http"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	resp := "200 OK"
	fmt.Fprint(w, resp)
}

func main() {
	// Retriving port for serving web server
	if (len(os.Args) != 2) {
		fmt.Printf("Please specify the port for web server.\n")
		os.Exit(1)
	} else {
	portArg, _ := strconv.Atoi(os.Args[1])
	
	if (portArg >= 1 && portArg <= 65535) {
		port := ":" + strconv.Itoa(portArg)
		
		// Route: /
		http.HandleFunc("/", root_handler)
		log.Fatal(http.ListenAndServe(port, nil))
	
	} else {
		fmt.Printf("The port you specified is invalid.\n")
		os.Exit(1)
	}
	}
}