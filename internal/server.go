package internal

import (
	"log"
	"net/http"
	"URIShorter/internal/controller"

)

func Run() {

	server := controller.SetUpRouter()
	defer server.DB.DB.Close()
	log.Println("Server listening on port 3000 ")
	http.ListenAndServe(":3000", server.Router)
	log.Fatal("Server is closed")

}
