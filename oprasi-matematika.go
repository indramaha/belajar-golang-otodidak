package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var hasil = 10 * 10
	fmt.Println(hasil)

	var g = 10
	g += 10 // a = a + 10
	fmt.Println(g)

	g++ // g = g + 1
	fmt.Println(g)

	var negatif = -100
	var positif = +100
	fmt.Println(negatif)
	fmt.Println(positif)
}
