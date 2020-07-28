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
	res, err := DoRequest(testServer, http.MethodGet, "/greet/ajitem")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message": "Hello, ajitem!"}`, string(body))
}

func TestFibonacci(t *testing.T) {
	// should return {"number": <n>, "fibonacci": <nth fibonacci number>} where the number n
	// is given by the user.
	var fibotests = []struct {
		name   string
		input  string
		output string
		err    bool
	}{
		{name: "valid input 22", input: "22", output: "17711"},
		{name: "invalid input string", input: "abcd", err: true},
	}

	for _, ft := range fibotests {
		t.Run(ft.name, func(t *testing.T) {
			res, err := DoRequest(testServer, http.MethodGet, fmt.Sprintf("/fibonacci/%s", ft.input))
			assert.NoError(t, err)

			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err)

			if ft.err {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Equal(t, `{"error": "fibonacci endpoint accepts only numbers"}`, string(body))
			} else {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, fmt.Sprintf(`{"number": %s, "fibonacci": %s}`, ft.input, ft.output), string(body))
			}
		})
	}
}
