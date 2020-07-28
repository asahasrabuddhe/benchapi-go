package bench_api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
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

func ParseBody(body io.Reader) *response {
	var res response

	decoder := json.NewDecoder(body)
	_ = decoder.Decode(&res)

	return &res
}

func TestMain(m *testing.M) {
	testServer = NewServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	// should return {"message": "Hello, world!"}
	res, err := DoRequest(testServer, http.MethodGet, "/")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)

	resp := ParseBody(res.Body)
	assert.Equal(t, "Hello, world!", resp.Message)
}

func TestGreet(t *testing.T) {
	// should return {"message": "Hello, <name>!"} where the name is given by the user
	var tests = []struct {
		name string
		err  bool
	}{
		{name: "ajitem"},
		{name: "radha"},
		{name: "", err: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("greet %s", tt.name), func(t *testing.T) {
			res, err := DoRequest(testServer, http.MethodGet, fmt.Sprintf("/greet/%s", tt.name))
			assert.NoError(t, err)

			resp := ParseBody(res.Body)
			if tt.err {
				assert.Equal(t, http.StatusNotFound, res.Code)
			} else {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, fmt.Sprintf("Hello, %s!", tt.name), resp.Message)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
	var tests = []struct {
		name      string
		input     string
		fibonacci string
		err       bool
	}{
		{name: "valid input 1", input: "1", fibonacci: "1"},
		{name: "valid input 2", input: "2", fibonacci: "1"},
		{name: "valid input 3", input: "3", fibonacci: "2"},
		{name: "valid input 22", input: "22", fibonacci: "17711"},
		{name: "valid input 32", input: "32", fibonacci: "2178309"},
		{name: "valid input 64", input: "64", fibonacci: "10610209857723"},
		{name: "invalid input string", input: "abcd", err: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DoRequest(testServer, http.MethodGet, fmt.Sprintf("/fibonacci/%s", tt.input))
			assert.NoError(t, err)

			resp := ParseBody(res.Body)
			if tt.err {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Equal(t, "fibonacci endpoint accepts only numbers", resp.Error)
			} else {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, tt.input, resp.Input)
				assert.Equal(t, tt.fibonacci, resp.Fibonacci)
			}
		})
	}
}
