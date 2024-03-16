package internal

import (
	"log"
	"net/http"
)

func Run() {

	server := setUpRouter()
	defer server.db.DB.Close()
	log.Println("Server listening on port 3000 ")
	http.ListenAndServe(":3000", server.router)
	log.Fatal("Server is closed")

}
