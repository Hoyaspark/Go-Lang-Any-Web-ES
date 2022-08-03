package main

import (
	"crypto/tls"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"strconv"
)

type Serv struct {
	*http.Server
	*chi.Mux
}

func NewServ(port int, tls *tls.Config) *Serv {

	mux := chi.NewRouter()

	return &Serv{
		&http.Server{
			Addr:         ":" + strconv.Itoa(port),
			Handler:      mux,
			TLSConfig:    tls,
			ReadTimeout:  0,
			WriteTimeout: 0,
			IdleTimeout:  0,
		},
		mux,
	}
}

func (s *Serv) SetMiddlewares() {
	s.Use(middleware.Logger)
	s.Use(middleware.Recoverer)
	s.Use(middleware.CleanPath)
	s.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}).Handler)
}

func (s *Serv) Start() error {
	if s.TLSConfig != nil {
		return s.ListenAndServeTLS("", "")
	}

	return s.ListenAndServe()
}
