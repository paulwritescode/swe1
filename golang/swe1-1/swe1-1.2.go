package main

import "fmt"

func Swe2() {
	names := []string{"Mairion", "Chepkoech", "Tuwei"}
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
}
