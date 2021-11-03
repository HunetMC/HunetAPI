package web

import (
	// own packages
	"github.com/HunetMC/HunetAPI/query"
	"github.com/HunetMC/HunetAPI/util"

	// go get packages
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	resp := "200 OK"
	fmt.Fprint(w, resp)
}

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	resp := query.GetPlayers()
	fmt.Fprint(w, resp)
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/player")
    _, uuid := filepath.Split(sub)
    if uuid != "" {
        if util.IsValidUUID(uuid) {
            resp := query.GetPlayer(uuid)
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