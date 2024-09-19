# Worker Pool

## Overview

This Go example demonstrates how to use worker pool using goroutines in a better way. This handles resource-intensive operations by utilizing multiple workers, monitors system resources (CPU and memory), and adjusts processing to ensure the system doesn't become overloaded. The application uses Goroutines and channels to distribute jobs (reports) across multiple workers.

## Features

- **Worker Pool**: A pool of Goroutines that process jobs concurrently.
- **CPU & Memory Monitoring**: The system's CPU and memory usage are monitored before each job is processed, and throttling occurs when resource usage exceeds a certain threshold.
- **Efficient Resource Management**: Memory and CPU usage are controlled by dynamically adjusting the job processing rate.
