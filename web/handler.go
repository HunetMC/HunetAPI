package web

import (
	// own packages
	"github.com/HunetMC/HunetAPI/database"
	"github.com/HunetMC/HunetAPI/util"
	
	// go get packages
	"fmt"
	"strings"
	"path/filepath"
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

func player_handler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/player")
    _, uuid := filepath.Split(sub)
    if uuid != "" {
        if (util.IsValidUUID(uuid)) {
            resp := database.Get_player(uuid)
			fmt.Fprint(w, resp)
        } else {
			resp := "Specified UUID is invalid."
			fmt.Fprint(w, resp)
        }
    } else {
        resp := "Please specify UUID."
        fmt.Fprint(w, resp)
    }
}