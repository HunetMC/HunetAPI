package web

import (
	// own packages
	"github.com/HunetMC/HunetAPI/database"
	
	// go get packages
	"fmt"
	"net/http"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	resp := "200 OK"
	fmt.Fprint(w, resp)
}

func players_handler(w http.ResponseWriter, r *http.Request) {
	
	resp := database.Get_players()
	fmt.Fprint(w, resp)
}