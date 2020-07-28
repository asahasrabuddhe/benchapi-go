# Bench API Go

The Bench API is a simple server written in Go to help benchmark the performance of the Go language. It has 3 endpoints

|Endpoint|Output|
|------|------|
|GET /|{"message": "Hello, world!"}|
|GET /greet/{name}|{"message": "Hello, <name>!"}|
|GET /fibonacci/{n}|{"number": n, "fibonacci": "<nth number in fibonacci series>}|

## Build

`make build` will build the `benchapi` docker image. To push the image to docker hub, use `make push`.

## Running the Server

`docker run --rm -p 8080:8080 ajitemsahasrabuddhe/benchapi:latest`
