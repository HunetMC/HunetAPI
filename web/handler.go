package web

import (
	"fmt"
	"net/http"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	resp := "200 OK"
	fmt.Fprint(w, resp)
}