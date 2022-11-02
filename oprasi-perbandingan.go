package main

import "fmt"

func main() {
	var name1 string = "bagus"
	var name2 string = "bagus"

	var hasil bool = name1 == name2
	fmt.Println(hasil)

	var value1 = 123
	var value2 = 234

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
