package main

import (
	"github.com/HunetMC/HunetAPI/database"
	"github.com/HunetMC/HunetAPI/web"
)

func main() {
	database.Open()
	web.Server()
}