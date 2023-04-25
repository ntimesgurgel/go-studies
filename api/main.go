package main

import (
	"api/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":80", nil)
}
