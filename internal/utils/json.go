package utils


import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
if code > 499{
	log.Println("Responding with 500 Error:" , msg)
}
/*
above code will allow me to to format my json like this
{
	error: "msg"
}
*/
	type errResponse struct{
		Error string `json:"error"`
	}
	RespondWithJSON(w,code,errResponse{
		Error:msg,
	})
}


func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)

}