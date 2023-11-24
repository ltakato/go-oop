package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Valor do saque inválido ou saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

// variadic function
func SomarVariosNumeros(numeros ...int) int {
	resultado := 0

	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	// comparison()

	// fmt.Println("Resultado SomarVariosNumeros:", SomarVariosNumeros(1, 2, 3, 4, 5))

	contaDoAgumon := ContaCorrente{titular: "Agumon", saldo: 500}
	contaDoGarurumon := ContaCorrente{titular: "Garurumon", saldo: 800}

	status := contaDoAgumon.Transferir(200, &contaDoGarurumon)

	fmt.Println("Status transferencia:", status)
	fmt.Println(contaDoAgumon)
	fmt.Println(contaDoGarurumon)
}

func comparison() {
	contaDoTakatittos := ContaCorrente{
		titular:       "Takatittos",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	// fmt.Println(contaDoTakatittos)

	contaDoTakatittos2 := ContaCorrente{
		titular:       "Takatittos",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	// go knows how to compare structs! (will not compare addresses for example)
	// will compare the content itself
	fmt.Println(contaDoTakatittos == contaDoTakatittos2)

	// contaDaBruna := ContaCorrente{
	// 	"Bruna", 222, 111222, 200,
	// }

	// fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"

	// will compare address!!
	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)

	// will compare content
	fmt.Println(*contaDaCris)
	fmt.Println(*contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}
