package swe1_4

import (
	"fmt"
	"sync"
)

/*
- concurrent greeter
- take a slice of names as input,
- launch a go routine fora each name to pring a greeting
- use sync.waitgroup to ensure all go routines completes before the main function exits
*/

func SayHello(sliceOfName []string){
	var wg sync.WaitGroup

	for index,name:= range sliceOfName{
		wg.Add(1)

		 go func(name string, index int){
			 printNames(name, index)
			 defer wg.Done() // decrement the counter when go routine completes
			}(name, index)

	}
	wg.Wait() //wait for all go routines to finish
}


func printNames(name string, index int) {
	fmt.Printf("%d: %s\n",index,name)
}