package main

import (
	"gin-rest/database"
	"gin-rest/routes"
)

func main() {
	database.ConnectionDB()
	routes.HandleRequest()
}
