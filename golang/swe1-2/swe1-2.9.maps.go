package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main(){
	const explanation string = "key: value pair"

//	one way of doing it
	var ages map[string]int
	ages = make(map[string]int)
	ages["paul"] = 30
	ages["Doe"] = 32

	fmt.Println(ages)

//	another way of doing it
emails := map[string]string{
	"john":"John@gmail.com",
	"Paul":"paul@gmail.com",
}
slug,status := json.MarshalIndent(emails, "", "  ")

if status != nil {
	log.Fatal("not working")
}

fmt.Println(string(slug))

}