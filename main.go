package main

import (
	"fmt"

	"github.com/ltakato/go-oop/contas/contas"
)

// variadic function
func SomarVariosNumeros(numeros ...int) int {
	resultado := 0

	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}

func main() {
	// comparison()

	// fmt.Println("Resultado SomarVariosNumeros:", SomarVariosNumeros(1, 2, 3, 4, 5))

	contaDoAgumon := contas.ContaCorrente{Titular: "Agumon", Saldo: 500}
	contaDoGarurumon := contas.ContaCorrente{Titular: "Garurumon", Saldo: 800}

	status := contaDoAgumon.Transferir(200, &contaDoGarurumon)

	fmt.Println("Status transferencia:", status)
	fmt.Println(contaDoAgumon)
	fmt.Println(contaDoGarurumon)
}

func comparison() {
	contaDoTakatittos := contas.ContaCorrente{
		Titular:       "Takatittos",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         125.5,
	}

	// fmt.Println(contaDoTakatittos)

	contaDoTakatittos2 := contas.ContaCorrente{
		Titular:       "Takatittos",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         125.5,
	}

	// go knows how to compare structs! (will not compare addresses for example)
	// will compare the content itself
	fmt.Println(contaDoTakatittos == contaDoTakatittos2)

	// contaDaBruna := ContaCorrente{
	// 	"Bruna", 222, 111222, 200,
	// }

	// fmt.Println(contaDaBruna)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = "Cris"

	// will compare address!!
	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)

	// will compare content
	fmt.Println(*contaDaCris)
	fmt.Println(*contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}
