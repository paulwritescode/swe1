package main

import "fmt"

func main(){
	const explain string = "Slice is a dynamically-sized view into the elemts of an array. They are much more common that array"

	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 3)

	fmt.Println(numbers)

	names := []string{"marion", "Praise", "Paul"}
	names2 := []rune("glad")

	fmt.Println(names)
	fmt.Println(names2)
	fmt.Println(string(names2))

}