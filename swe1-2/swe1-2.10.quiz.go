package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title string
	Aurthor string
	Price float64
}

func main() {
	const explanation string = "tests : variable declaration, array and slice, struct, map"

//	variable declaration
	var productName string = "Awesome Gadget"
	price := 99.99
	fmt.Println(productName, price)

//	Array and slice
	numbers:=[5]int{5, 10, 15, 20, 25}
	fmt.Println(numbers)

	numberSlice := numbers[1:4]
	fmt.Println(numberSlice)

//	struct
	book:= Book{Title: "Aurora", Aurthor: "luna", Price: 34.23}
	fmt.Println(book)

//	map
	products:=map[string]float64{
		"Book":34.23,
		"Pen":12.90,
		"Rubber":1.3,
	}

//	fmt.Println(products)

//	json marshal with indentation

	slug, ok := json.MarshalIndent(products, "", "  ")
	if ok != nil {
		log.Fatal("Marsal not working")
	}
	fmt.Println(string(slug))
}