package swe1_4

import (
	"fmt"
	"net/http"
)

func fetchUrl (url string, ch chan string){
resp, err := http.Get(url)

if err != nil {
	ch <- fmt.Sprintf("Error fetching %s: %s", url, err)
}
defer resp.Body.Close()
ch <- fmt.Sprintf("success fetched %s with status code: %d", url, resp.StatusCode)

}

func RoutinCall (){
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.golang.or",
	}

	ch := make(chan string) // create a channel to receiver results

	for _, url := range urls {
		go fetchUrl(url, ch) // creates a go routine for each url
	}

	// receiver results from the channel
	for i := 0; i<len(urls);i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("All urls fetched")
}