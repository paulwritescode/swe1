package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main (){
	const explanation string = "a composite meaning it is a creation of many thant one data type. It groups together many data types"

	var person Person
	person.Name = "paul"
	person.Age = 24

	fmt.Println(person)

	person2:= Person{Name:"mbugua", Age:45}
	fmt.Println(person2)
}