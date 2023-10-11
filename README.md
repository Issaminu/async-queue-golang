# async-queue-golang
A project where I'm exploring the basics of concurrency in Go. It shows how to implement an asynchronous job queue with a worker pool, focusing on concurrency while handling HTTP requests.
## Overview
This project showcases the following key components:

* A **Queue** data structure for managing a **list of jobs**.
* A **worker pool** of a specified size for **concurrently processing jobs**.
* Handling **HTTP requests** and processing them asynchronously.
## Getting Started
### Prerequisites
Before you can run the project, make sure you have Go installed on your system.

### Installation
To set up and run the project, follow these steps:

1. Clone the repository:
```shell
git clone https://github.com/your-username/async-queue-golang.git
cd async-queue-golang
```
2. Build and run the project:
```shell
go build
./async-queue-golang
```
The project will start a server and listen on port 8080.

## Usage
### Adding Jobs
To add a job to the queue, send an HTTP request to the `/addJob` endpoint. You can include a `Name` parameter in the request query, which will be used to simulate work processing.

Example:
```shell
curl -X GET 'http://localhost:8080/addJob?Name=YourName'
```

### Worker Pool
The project includes a worker pool of size `WORKER_POOL_SIZE` (set to 4 by default) for concurrently processing jobs from the queue. The workers will extract the `Name` parameter from the request query and send a response.

### Example Responses
If the "Name" parameter is provided:
```shell
Hello, YourName
```
If the "Name" parameter is missing or empty:
```shell
Error: Name parameter is missing or empty
```
