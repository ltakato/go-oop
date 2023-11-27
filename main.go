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

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	// comparison()

	// fmt.Println("Resultado SomarVariosNumeros:", SomarVariosNumeros(1, 2, 3, 4, 5))

	clienteAgumon := clientes.Titular{Nome: "Agumon", CPF: "123.111.123-12", Profissao: "Digimon"}
	contaDoAgumon := contas.ContaCorrente{Titular: clienteAgumon}
	contaDoAgumon.Depositar(500)

	PagarBoleto(&contaDoAgumon, 60)

	fmt.Println("Valor pós pagamento do boleto do Agumon: ", contaDoAgumon.ObterSaldo())

	clienteGarurumon := clientes.Titular{Nome: "Garurumon", CPF: "223.311.123-85", Profissao: "Digimon"}
	contaDoGarurumon := contas.ContaCorrente{Titular: clienteGarurumon}
	contaDoGarurumon.Depositar(800)

	status := contaDoAgumon.Transferir(200, &contaDoGarurumon)

	fmt.Println("Status transferencia:", status)
	fmt.Println(contaDoAgumon)
	fmt.Println("Saldo Agumon: ", contaDoAgumon.ObterSaldo())
	fmt.Println(contaDoGarurumon)
	fmt.Println("Saldo Garurumon: ", contaDoGarurumon.ObterSaldo())

	clientePatamon := clientes.Titular{Nome: "Patamon", CPF: "728.431.123-98", Profissao: "Digimon"}
	contaDoPatamon := contas.ContaPoupanca{Titular: clientePatamon}
	contaDoPatamon.Depositar(100)

	fmt.Println("Saldo do Patamon:", contaDoPatamon.ObterSaldo())
}

func comparison() {
	clienteTakatittos := clientes.Titular{
		Nome: "Takatittos",
	}
	contaDoTakatittos := contas.ContaCorrente{
		Titular:       clienteTakatittos,
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}
	contaDoTakatittos.Depositar(125.5)

	// fmt.Println(contaDoTakatittos)

	contaDoTakatittos2 := contas.ContaCorrente{
		Titular:       clienteTakatittos,
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}
	contaDoTakatittos2.Depositar(125.5)

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
