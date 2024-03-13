package internal

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	Router *chi.Mux
}

func Run() {

	router := setUpRouter()
	ser := Server{Router: router}

	http.ListenAndServe(":3000", ser.Router)
}
