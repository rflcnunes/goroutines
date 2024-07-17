package main

import (
	"fmt"
	"github.com/fatih/color"
	"goroutines/utils"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func workerFunc(workerID int, iterations int, maxSleep int) {
	red := color.New(color.FgRed).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	for i := 0; i < iterations; i++ {
		start := time.Now()
		for j := 0; j < 1000000; j++ {
			_ = rand.Intn(1000) * rand.Intn(1000)
		}
		duration := time.Since(start)

		if workerID == 10 {
			if i%2 == 0 {
				red("Worker %d: Iteration %d, Time of processing: %v\n", workerID, i, duration)
			} else {
				green("Worker %d: Iteration %d, Time of processing: %v\n", workerID, i, duration)
			}
		}

		time.Sleep(time.Duration(rand.Intn(maxSleep)) * time.Millisecond)
	}
	wg.Done()
}

func getInput(prompt string) int {
	fmt.Print(prompt)
	var input string
	_, _ = fmt.Scanln(&input)
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Error converting input to integer: %v\n", err)
	}
	return result
}

func main() {
	utils.SetupLogging()
	yellow := color.New(color.FgYellow).PrintfFunc()

	numCPUs := getInput("Type the number of CPUs to use: ")
	if numCPUs > runtime.NumCPU() {
		fmt.Printf("Max number of CPUs is %d\n", runtime.NumCPU())
		numCPUs = runtime.NumCPU()

		yellow("Automatically set to %d CPUs\n", numCPUs)
	}

	runtime.GOMAXPROCS(numCPUs)

	numWorkers := getInput("Type the number of workers: ")

	runtime.GOMAXPROCS(numCPUs)

	startTime := time.Now()

	wg.Add(numWorkers)

	iterations := 50
	maxSleep := 100

	for workerID := 1; workerID <= numWorkers; workerID++ {
		go workerFunc(workerID, iterations, maxSleep)
	}

	numberOfGoroutines := runtime.NumGoroutine()

	wg.Wait()

	executionTime := time.Since(startTime).Seconds()

	log.Printf("CPUs: %d, Goroutines: %d, Time: %.0fs, Workers: %d\n", runtime.GOMAXPROCS(0), numberOfGoroutines, executionTime, numWorkers)
}
