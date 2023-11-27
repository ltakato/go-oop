package main

import (
	"fmt"

	"github.com/ltakato/go-oop/clientes"
	"github.com/ltakato/go-oop/contas"
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

	clienteAgumon := clientes.Titular{Nome: "Agumon", CPF: "123.111.123-12", Profissao: "Digimon"}
	contaDoAgumon := contas.ContaCorrente{Titular: clienteAgumon, Saldo: 500}
	clienteGarurumon := clientes.Titular{Nome: "Garurumon", CPF: "223.311.123-85", Profissao: "Digimon"}
	contaDoGarurumon := contas.ContaCorrente{Titular: clienteGarurumon, Saldo: 800}

	status := contaDoAgumon.Transferir(200, &contaDoGarurumon)

	fmt.Println("Status transferencia:", status)
	fmt.Println(contaDoAgumon)
	fmt.Println(contaDoGarurumon)
}

func comparison() {
	clienteTakatittos := clientes.Titular{
		Nome: "Takatittos",
	}
	contaDoTakatittos := contas.ContaCorrente{
		Titular:       clienteTakatittos,
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         125.5,
	}

	// fmt.Println(contaDoTakatittos)

	contaDoTakatittos2 := contas.ContaCorrente{
		Titular:       clienteTakatittos,
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
	contaDaCris.Titular = clientes.Titular{Nome: "Cris"}

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = clientes.Titular{Nome: "Cris"}

	// will compare address!!
	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)

	// will compare content
	fmt.Println(*contaDaCris)
	fmt.Println(*contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}
