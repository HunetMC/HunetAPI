package web

import (
	"fmt"
	"net/http"
)

func Server() {
	// Routes
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/player/", player_handler)
	http.HandleFunc("/players", players_handler)
	
	err := http.ListenAndServe(get_port(), nil)
	if err != nil {
		fmt.Println("Failed to start web server.")
		return
	} else {
		fmt.Println("Successfully started web server.")
	}
}