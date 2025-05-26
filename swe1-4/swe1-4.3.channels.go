package swe1_4

import "fmt"

func SendData(ch chan<- int){
	ch<- 35
}

func ReceiveData(ch <-chan int){
	value:= <-ch
	fmt.Println("Running from channel",value)
}