package swe1_4

import "fmt"

func CloseChannel(){
	ch := make(chan int ,4)
	ch<-1
	ch<-2
	ch<-3
	ch<-4
	close(ch) // we close the channel

	for value := range ch {
		fmt.Println("from the closed channel",value)
	}

	val, ok := <-ch
	fmt.Println(val, ok)


}