package bench_api

import (
	"fmt"
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
	testServer = NewServer()
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

			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err)
			if tt.err {
				assert.Equal(t, http.StatusNotFound, res.Code)
			} else {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, fmt.Sprintf(`{"message": "Hello, %s!"}`, tt.name), string(body))
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
	var tests = []struct {
		name   string
		input  string
		output string
		err    bool
	}{
		{name: "valid input 1", input: "1", output: "1"},
		{name: "valid input 2", input: "2", output: "1"},
		{name: "valid input 3", input: "3", output: "2"},
		{name: "valid input 22", input: "22", output: "17711"},
		{name: "valid input 32", input: "32", output: "2178309"},
		{name: "valid input 64", input: "64", output: "10610209857723"},
		{name: "invalid input string", input: "abcd", err: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DoRequest(testServer, http.MethodGet, fmt.Sprintf("/fibonacci/%s", tt.input))
			assert.NoError(t, err)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err)

			if tt.err {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Equal(t, `{"error": "fibonacci endpoint accepts only numbers"}`, string(body))
			} else {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, fmt.Sprintf(`{"number": %s, "fibonacci": %s}`, tt.input, tt.output), string(body))
			}
		})
	}
}
