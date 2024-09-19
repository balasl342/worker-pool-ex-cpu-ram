# Worker Pool

## Overview

This Go example demonstrates how to use worker pool using goroutines in a better way. This handles resource-intensive operations by utilizing multiple workers, monitors system resources (CPU and memory), and adjusts processing to ensure the system doesn't become overloaded. The application uses Goroutines and channels to distribute jobs (reports) across multiple workers.

## Features

- **Worker Pool**: A pool of Goroutines that process jobs concurrently.
- **CPU & Memory Monitoring**: The system's CPU and memory usage are monitored before each job is processed, and throttling occurs when resource usage exceeds a certain threshold.
- **Efficient Resource Management**: Memory and CPU usage are controlled by dynamically adjusting the job processing rate.

## How It Works

1. **Worker Pool**:
   - The worker pool consists of a specified number of workers (default is 10).
   - Each worker picks up a job from the queue (reports to be generated) and processes it concurrently.

2. **Job Distribution**:
   - The jobs (report generation tasks) are sent to workers through a buffered channel to avoid blocking.
   - Workers continue processing until the job queue is empty.

3. **Resource Monitoring**:
   - Before processing each job, the application checks the system's current CPU and memory usage.
   - If resource usage exceeds 80%, the processing is throttled by introducing a short sleep.

4. **Report Generation**:
   - The core business logic for generating a report can be defined inside the `generateReport` function.

## Code Breakdown

### `main.go`

- **`UserController`**: A placeholder struct representing the user management module.
  
- **`generateReport`**: The function responsible for generating a report. It takes in a `UserController` instance and a job (`ps`).

- **`workerPool`**: The function that implements the worker pool. Each worker processes jobs from the `jobs` channel and calls the `generateReport` function.

- **`monitorResourceUsage`**: This function checks the current system's CPU and memory usage. If either exceeds 80%, it introduces a delay to prevent overloading.

- **`getCpuUsage`**: A simple function that returns CPU usage by calculating the ratio of Goroutines to CPU cores.

- **`downloadReport`**: The main function that initializes the worker pool and distributes jobs to workers.

- **`getData`**: A placeholder function that simulates fetching data for report generation. In the actual implementation, this would retrieve real data.

