package main

import (
	"fmt"
	"sync"
	"time"
)

func callDB(wg *sync.WaitGroup) {
	fmt.Println("Calling DB...")
	time.Sleep(1 * time.Second)
	fmt.Println("DB call done.")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	fmt.Println("Calling API...")
	time.Sleep(2 * time.Second)
	fmt.Println("API call done.")
	wg.Done()
}

func callExternalService(wg *sync.WaitGroup) {
	fmt.Println("Calling external service...")
	time.Sleep(3 * time.Second)
	fmt.Println("External service call done.")
	wg.Done()
}

// WaitGroup
func main() {
	fmt.Println("Hello, World!")

	var wg sync.WaitGroup
	wg.Add(3)
	go callDB(&wg)
	go callAPI(&wg)
	go callExternalService(&wg)

	wg.Wait()
}
