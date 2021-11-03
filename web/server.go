package web

import (
	"fmt"
	"net/http"
)

func Server() {
	// Routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/player/", PlayerHandler)
	http.HandleFunc("/players", PlayersHandler)
	http.HandleFunc("/duel/", DuelHandler)
	
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		fmt.Println("Failed to start web server.")
		return
	} else {
		fmt.Println("Successfully started web server.")
	}
}