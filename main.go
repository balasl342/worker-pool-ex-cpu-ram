package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type UserController struct{}

var wg sync.WaitGroup

// Function to generate the report
func generateReport(c *UserController, ps interface{}) {
	defer wg.Done()
	// Business logic here
}

// Function to create a worker pool to process the reports
func workerPool(c *UserController, jobs <-chan interface{}, workerID int) {
	for ps := range jobs {
		monitorResourceUsage() // Check the CPU and memory usage before processing each job
		generateReport(c, ps)
	}
}

// Monitor CPU and memory usage and throttle processing if needed
func monitorResourceUsage() {
	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// Get the percentage of memory being used
		memoryUsage := float64(memStats.Alloc) / float64(memStats.Sys) * 100

		// Get the number of CPU cores utilized
		cpuUsage := getCpuUsage()

		// Print the resource usage (for debugging)
		fmt.Printf("Memory Usage: %.2f%% | CPU Usage: %.2f%%\n", memoryUsage, cpuUsage)

		// Throttle if CPU or memory usage exceeds 80%
		if cpuUsage < 80 && memoryUsage < 80 {
			break // Continue processing if utilization is below 80%
		}

		// Throttle processing by sleeping for a short duration to allow system to recover
		time.Sleep(100 * time.Millisecond)
	}
}

// CPU usage check
func getCpuUsage() float64 {
	return float64(runtime.NumGoroutine()) / float64(runtime.NumCPU()) * 100
}

// Download report
func downloadReport(c *UserController) {
	runtime.GC() // trigger garbage collection

	noOfWorkers := 10
	batchData := getData() // replace with actual data retrieval function

	// Create a job queue to distribute work
	jobs := make(chan interface{}, len(batchData)) // buffered channel to prevent blocking

	// Start worker pool
	for w := 0; w < noOfWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			workerPool(c, jobs, workerID)
		}(w)
	}

	// Send jobs to workers
	for _, ps := range batchData {
		jobs <- ps
	}

	// Close the jobs channel to signal workers to stop
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
}

// Data retrieval
func getData() []interface{} {
	// return batch data (simulated array of records)
	return make([]interface{}, 20000)
}

func main() {
	c := &UserController{}
	downloadReport(c)
}
