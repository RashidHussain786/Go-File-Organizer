package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type JobResult struct {
	URL    string
	Result string
	Err    error
}

func main() {
	urls, err := readFile("urls.txt")
	if err != nil {
		return
	}
	if len(urls) == 0 {
		println("No URLs found in the file.")
		return
	}
	// workerPattern(urls)
	// consumerPattern(urls)
	robustWorkerPattern(urls)

	println("Total URLs:", len(urls))
	println("Done reading URLs from file.")
}

func readFile(filePath string) ([]string, error) {
	var urls []string
	files, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer files.Close()

	scanner := bufio.NewScanner(files)
	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			continue
		}
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}

func workerPattern(urls []string) {
	ch := make(chan string, 100)
	for _, url := range urls {
		go func(u string) {
			result := "Worker pattern URL: " + u
			ch <- result // Simulating some processing
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-ch // Receiving results from the channel
		println(result)
	}
	close(ch)
}

func consumerPattern(urls []string) {
	ch := make(chan string)
	var wg sync.WaitGroup

	// Consumer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for url := range ch {
			println("Consumer pattern URL:", url)
		}
	}()

	// Producer
	for _, url := range urls {
		ch <- url
	}
	close(ch) // Signal consumer to finish

	wg.Wait()
	println("All URLs processed.")
}

func robustWorkerPattern(urls []string) {
	chTasks := make(chan string, len(urls))
	chResults := make(chan JobResult, len(urls))

	numberWorkers := 10
	var wg sync.WaitGroup

	for i := 0; i < numberWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for url := range chTasks {
				processTask(url, workerID, chResults)
			}
		}(i)
	}

	for _, url := range urls {
		chTasks <- url
	}
	close(chTasks)

	wg.Wait()
	close(chResults)

	for result := range chResults {
		if result.Err != nil {
			fmt.Printf("Error processing URL %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Println(result.URL, result.Result)
		}
	}
}

func processTask(url string, workerID int, chResults chan<- JobResult) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("worker %d panicked while processing %s: %v", workerID, url, r)
			chResults <- JobResult{URL: url, Result: "", Err: err}
		}
	}()

	fmt.Printf("Worker %d started job: %s\n", workerID, url)
	// time.Sleep(time.Second)  // for simulating work
	resp, err := http.Get(url)
	if err != nil {
		chResults <- JobResult{URL: url, Result: "Down", Err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		status := fmt.Sprintf("UP - %s", resp.Status)
		chResults <- JobResult{URL: url, Result: status, Err: nil}
	} else {
		status := fmt.Sprintf("DOWN - %s", resp.Status)
		chResults <- JobResult{URL: url, Result: status, Err: nil}
	}
}
