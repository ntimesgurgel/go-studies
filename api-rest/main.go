package main

import (
	"api/rest/database"
	"api/rest/models"
	"api/rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor REST com GO")
	routes.HandleRequest()
}
