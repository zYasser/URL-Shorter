package controller

import (
	"URIShorter/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type url struct {
	URL string `json:"url"`
}

func HandlerURL( w http.ResponseWriter, r *http.Request) {
	var urlReq url
	err := json.NewDecoder(r.Body).Decode(&urlReq)
	if err != nil {
		utils.RespondWithError(w, 400, "There's No URL")
	}
	new_url := url{
		URL: utils.GenerateShortURL(),
	}
	
	fmt.Printf("%s new url \n" ,new_url )
	utils.RespondWithJSON(w,200,new_url)
}
