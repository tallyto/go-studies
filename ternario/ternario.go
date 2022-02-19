package main

import "fmt"

// Não existe operador ternário em Go.

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	fmt.Println(obterResultado(6.2))
	fmt.Println(obterResultado(4.9))
}
