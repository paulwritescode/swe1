package main

import "fmt"

type days int

const (
	sunday days = iota
	monday
	teusday
	wednesday
	thursday
	friday
	saturday
)

type Event struct {
	Name string
	Day days
}

func main(){
meetingDay := teusday
fmt.Println("meeting is on", meetingDay)

// using days in strunct
party:= Event{
	Name: "See my baby",
	Day:saturday,
}

fmt.Printf("%s is on %v\n", party.Name, party.Day)

fmt.Println("Values of the days constants:")
fmt.Println("Sunday:", sunday)
}