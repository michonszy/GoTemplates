package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Status int
}

func worker(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode}
	}
}
func main() {
	fmt.Println("Go workers")
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	urls := []string{
		"https://wp.pl",
		"https://onet.pl",
		"https://google.com",
		"https://youtube.pl",
		"https://wp.pl",
		"https://onet.pl",
		"https://google.com",
		"https://youtube.pl",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 8; a++ {
		result := <-results
		log.Println(result)
	}
}
