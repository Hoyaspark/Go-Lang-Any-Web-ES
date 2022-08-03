package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	serv := NewServ(8080, nil)

	serv.SetMiddlewares()

	apiRouter := chi.NewRouter()

	apiRouter.Group(func(r chi.Router) {
		r.Get("/", func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("test"))
			return
		})

		r.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("hello"))
			return
		})
	})

	serv.Mount("/api", apiRouter)

	if err := serv.Start(); err != nil {
		panic("cannot create web server")
	}
}
