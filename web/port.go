package web

import (
	"os"
	"fmt"
	"strconv"
)

func get_port() string {
	var port string
	
	// Retriving port for serving web server
	if (len(os.Args) != 2) {
		fmt.Printf("Please specify the port for web server.\n")
		os.Exit(1)
	} else {
	portArg, _ := strconv.Atoi(os.Args[1])
	
	if (portArg >= 1 && portArg <= 65535) {
		port = ":" + strconv.Itoa(portArg)
	} else {
		fmt.Printf("The port you specified is invalid.\n")
		os.Exit(1)
	}
	}
	
	return port
}