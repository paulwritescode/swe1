package main

import "fmt"

func main(){
	var numbers[5]int
	numbers[0] = 1
	numbers[2] = 4
	numbers[1] = 34
	numbers[3] = 10
	numbers[4] = 41

	names:=[2]string{"alice", "john"}

	fmt.Println(names)
	fmt.Println(numbers)
}