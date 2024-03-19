package controller

import (
	"URIShorter/internal/utils"
	"encoding/json"
	"net/http"
)

type url struct {
	URL string `json:"url"`
}

func (con *Config) handlerURL(w http.ResponseWriter, r *http.Request) {
	var urlReq url
	err := json.NewDecoder(r.Body).Decode(&urlReq)
	if err != nil {
		utils.RespondWithError(w, 400, "There's No URL")
	}
	new_url := url{
		URL: utils.GenerateShortURL(),
	}
	con.DB.InsertUrl(urlReq.URL, new_url.URL)
	utils.RespondWithJSON(w, 200, new_url)

}
