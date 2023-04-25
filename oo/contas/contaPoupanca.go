package contas

import "oo/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	Saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && c.Saldo > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "valor sacado"
	} else {
		return "valor insuficiente"
	}

}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "depósito realizado com sucesso", c.Saldo
	} else {
		return "valor para depósito negativo", c.Saldo
	}
}
