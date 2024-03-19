package controller

import (
	"log"
	"net/http"

	"URIShorter/internal/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Config struct {
	DB     model.Query
	Router *chi.Mux
}

func SetUpRouter() *Config {
	log.Println("Set up a router")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	con := &Config{DB: model.OpenDB(), Router: r}
	r.Post("/new_url", con.handlerURL)
	log.Println("Finished Set up router")

	return con

}
