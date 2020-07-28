package bench_api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type server struct {
	router *chi.Mux
}

func newServer() *server {
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

	s.router.Get("/{name}", func(writer http.ResponseWriter, request *http.Request) {
		name := chi.URLParam(request, "name")

		_, _ = fmt.Fprintf(writer, `{"message": "Hello, %s!"}`, name)
	})
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}
