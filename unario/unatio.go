package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x = x + 1 ou x += 1
	y-- // y = y - 1 ou y -= 1
	print(x, y)
	fmt.Println(x != y)
}
