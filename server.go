package bench_api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type server struct {
	router *chi.Mux
}

func NewServer() *server {
	s := &server{
		router: chi.NewRouter(),
	}

	s.routes()

	return s
}

func (s *server) routes() {
	s.router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, `{"message": "Hello, world!"}`)
	})

	s.router.Get("/greet/{name}", func(writer http.ResponseWriter, request *http.Request) {
		name := chi.URLParam(request, "name")

		_, _ = fmt.Fprintf(writer, `{"message": "Hello, %s!"}`, name)
	})

	s.router.Get("/fibonacci/{number}", func(writer http.ResponseWriter, request *http.Request) {
		numberStr := chi.URLParam(request, "number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(writer, `{"error": "fibonacci endpoint accepts only numbers"}`)
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

		_, _ = fmt.Fprintf(writer, `{"number": %d, "fibonacci": %d}`, number, result)
	})
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
