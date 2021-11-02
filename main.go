package main

import (
	// own packages
	"github.com/HunetMC/HunetAPI/db"
	"github.com/HunetMC/HunetAPI/web"
)

func main() {
	db.Connect()
	web.Server()
}