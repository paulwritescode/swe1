package swe1_4

import (
	"fmt"
	"time"
)

func greetings(name string){
fmt.Println("Hello routine:" +name)
}

func Myroutines(){
	go greetings("Alice")
	go greetings("doe")
	time.Sleep(1* time.Second)
	fmt.Println("Main function call after routine")
}