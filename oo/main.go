package main

import (
	"fmt"
	"oo/clientes"
	"oo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	conta := contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Nathan", Cpf: "000.000.000-00", Profissao: "Quebracodigos"},
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         500,
	}

	fmt.Println(conta.Titular)
}
