package swe1_4

import "time"

type Job struct {
	ID int
	RandomNumber int
}

type Result struct {
	JobID int
	SumofDigits int
	Name string
}

func digits(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)
	return sum
}

func Worker(jobs <-chan Job, results chan<- Result){
	for job := range jobs {
		output := Result{
			JobID: job.ID,
			SumofDigits: digits(job.RandomNumber),
			Name: "hello all",
		}
		results <- output
	}
}