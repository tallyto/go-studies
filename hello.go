package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14
	var raio = 3.2 // tipo inferido pelo compilador
	area := pi * math.Pow(raio, 2)
	fmt.Println("Valor area", area)
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)
	var e, f bool = true, false
	fmt.Println(e, f)
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
