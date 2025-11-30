package main

import (
	"net/http"
	"time"

	"github.com/bruneldev/ecom-api/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ponnnnng !!!"))
	})
	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	return r

	// http.ListenAndServe(":3333", r)
}

func (app *application) run(h http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	return srv.ListenAndServe()
}

type application struct {
	config config
}
type config struct {
	addr string
	db   dbConfig
}
type dbConfig struct{
	dsn string
}

func sm(){
	func(func())func(){return func(){}}(func(){});
}
