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

func main() {
	// comparison()

	contaDoAgumon := ContaCorrente{}
	contaDoAgumon.titular = "Agumon"
	contaDoAgumon.saldo = 500

	fmt.Println("Saldo do agumon (antes)", contaDoAgumon.saldo)

	fmt.Println(contaDoAgumon.Sacar(300))

	fmt.Println("Saldo do agumon (depois)", contaDoAgumon.saldo)

	fmt.Println(contaDoAgumon.Sacar(-100))

	fmt.Println("Saldo do agumon (depois)", contaDoAgumon.saldo)
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
