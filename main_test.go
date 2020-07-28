package bench_api

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	res := httptest.NewRecorder()
	testServer.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "Hello, world!"}`, string(body))
}

func TestGreet(t *testing.T) {
	// should return {"message": "Hello, <name>!"} where the name is given by the user
	req, err := http.NewRequest(http.MethodGet, "/ajitem", nil)
	assert.NoError(t, err)

	res := httptest.NewRecorder()
	testServer.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "Hello, ajitem!"}`, string(body))
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
}