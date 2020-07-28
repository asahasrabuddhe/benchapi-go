package bench_api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strconv"
)

type server struct {
	router *chi.Mux
}

type response struct {
	Message   string `json:"message,omitempty"`
	Input     string `json:"input,omitempty"`
	Fibonacci string `json:"fibonacci,omitempty"`
	Error     string `json:"error,omitempty"`
}

func NewServer() *server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := &server{
		router: r,
	}

	s.routes()

	return s
}

func (s *server) routes() {
	s.router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		res := &response{Message: "Hello, world!"}

		encoder := json.NewEncoder(writer)
		_ = encoder.Encode(res)
	})

	s.router.Get("/greet/{name}", func(writer http.ResponseWriter, request *http.Request) {
		name := chi.URLParam(request, "name")
		res := &response{Message: fmt.Sprintf("Hello, %s!", name)}

		encoder := json.NewEncoder(writer)
		_ = encoder.Encode(&res)
	})

	s.router.Get("/fibonacci/{number}", func(writer http.ResponseWriter, request *http.Request) {
		encoder := json.NewEncoder(writer)

		numberStr := chi.URLParam(request, "number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			res := &response{Error: "fibonacci endpoint accepts only numbers"}
			_ = encoder.Encode(&res)

			return
		}

		fibonacci := func() func() int64 {
			a, b := int64(0), int64(1)
			return func() int64 {
				defer func() {
					a, b = b, a+b
				}()
				return a
			}
		}()

		var result int64
		for i := 0; i <= number; i++ {
			result = fibonacci()
		}

		res := &response{Input: fmt.Sprintf("%d", number), Fibonacci: fmt.Sprintf("%d", result)}
		_ = encoder.Encode(&res)
	})
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
