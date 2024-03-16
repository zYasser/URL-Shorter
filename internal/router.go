package internal

import (
	"log"
	"net/http"

	"URIShorter/internal/controller"
	"URIShorter/internal/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type config struct {
	db     model.Query
	router *chi.Mux
}

func setUpRouter() *config {
	log.Println("Set up a router")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/new_url", controller.HandlerURL)
	con := &config{db: model.OpenDB(), router: r}
	log.Println("Finished Set up router")

	return con

}
