/*
- uncomment for swe1-1 - swe1-3
*/

//package main
//
//import (
//    "time"
////    swe1_3 "github.com/paulwritescode/swe1/swe1-3"
////    swe1_4 "github.com/paulwritescode/swe1/swe1-4"
//    "fmt"
//    swe1_4 "github.com/paulwritescode/swe1/swe1-4"
//)
//
//func main (){
////    fmt.Println("Hello world from main file")
////    time.Sleep(1* time.Second)
////    swe1_3.Myfunction()
////
////    time.Sleep(1* time.Second)
////    gratitude:=swe1_3.Packages()
////    fmt.Println(gratitude)
////
////    time.Sleep(1* time.Second)
////    reversedString := swe1_3.ReverseString("mariontuwei")
////    fmt.Println(reversedString)
////
////    time.Sleep(1* time.Second)
////    swe1_4.Myroutines()
////    time.Sleep(1* time.Second)
////
////    swe1_4.RoutinCall()
//ch:= make(chan int)
//    go swe1_4.SendData(ch)
//    go swe1_4.ReceiveData(ch)
//    time.Sleep(1* time.Second)
//
//fmt.Println("This is after channesl and go routines")
//swe1_4.CloseChannel()
//
//}


/*
- uncomment for swe1-4
*/

//package main
//
//import (
//    "fmt"
//    swe1_4 "github.com/paulwritescode/swe1/swe1-4"
//    "math/rand"
//)
//
//func main(){
//    noOfJobs := 10
//    noOfWorkers := 3
//    var jobs = make(chan swe1_4.Job, noOfJobs)
//    var results = make(chan swe1_4.Result, noOfJobs)
//
//    // create worker pool
//    for i := 0; i < noOfWorkers; i++ {
//        go swe1_4.Worker(jobs, results)
//    }
//
//    // allocate jobs
//    for i := 0; i <noOfJobs; i++ {
//        randomNo := rand.Intn(999)
//        job := swe1_4.Job{ID:i, RandomNumber: randomNo}
//        jobs <- job
//    }
//
//    close(jobs)
//    // collect results
//    for i := 0; i < noOfJobs; i++{
//        result := <- results
//        jobi := <- jobs
//        fmt.Printf("Worker: %s Job id %d, input random no %d , sum of digits %d\n",result.Name, result.JobID, jobi.RandomNumber, result.SumofDigits)
//    }
//}



/*
- uncomment for swe1-4 - quizes
*/

package main

import swe1_4 "github.com/paulwritescode/swe1/swe1-4"

func main(){
    names := []string{"marion", "dawn", "king"}
    swe1_4.SayHello(names)
}