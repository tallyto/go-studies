package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 3)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Print("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{"João", 20}
	p2 := Pessoa{"João", 20}

	fmt.Println("Pessoas:", p1 == p2)

	print(p1.Nome)
	print(p2.Idade)
}
