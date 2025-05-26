package swe1_3

import "fmt"

var initialized bool

func init(){
	fmt.Println("We are initializing the function at swe1-3.2.initialization .......")
	initialized = true
}

func Myfunction(){
	if initialized {
		fmt.Println("My function at swe1-3.2.initialization is ready to use : 22")
	}
}