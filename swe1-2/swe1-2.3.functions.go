package main

import "fmt"

func main(){
	const name = "Paul Kinyatti"
	const year = 2001

//	year:=2003
// above we get the error that we cannot reasign a constant

	fmt.Println(name,year +1)
}