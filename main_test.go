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

func DoRequest(handler http.Handler, method, url string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	return res, nil
}

func TestMain(m *testing.M) {
	testServer = newServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	// should return {"message": "Hello, world!"}
	res, err := DoRequest(testServer, http.MethodGet, "/")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "Hello, world!"}`, string(body))
}

func TestGreet(t *testing.T) {
	// should return {"message": "Hello, <name>!"} where the name is given by the user
	res, err := DoRequest(testServer, http.MethodGet, "/ajitem")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "Hello, ajitem!"}`, string(body))
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
}