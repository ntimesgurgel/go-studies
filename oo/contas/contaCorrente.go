package contas

import "oo/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	Saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && c.Saldo > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "valor sacado"
	} else {
		return "valor insuficiente"
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "depósito realizado com sucesso", c.Saldo
	} else {
		return "valor para depósito negativo", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, conta *ContaCorrente) string {
	resultado := c.Sacar(valorDaTransferencia)

	if resultado == "valor sacado" {
		conta.Depositar(valorDaTransferencia)
		return "transferência realizada com sucesso"
	}

	return resultado
}
