package bench_api

import (
	"os"
	"testing"
)

var testServer *server

func TestMain(m *testing.M) {
	testServer = newServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	// should return {"message": "Hello, world!"}
}

func TestGreet(t *testing.T) {
	// should return {"message": "Hello, <name>!"} where the name is given by the user
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
}