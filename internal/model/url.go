package model

import (
	"fmt"
	"time"
)

type URL struct {
	url        string
	id         string
	created_at time.Time
}

func (query *Query) InsertUrl(url string, hashUrl string) {
	var newUrl URL
	_ = newUrl
	result := query.DB.QueryRow("INSERT INTO urls (id, url) VALUES ($1, $2) RETURNING *", hashUrl, url)

	err := result.Scan(&newUrl.created_at, &newUrl.id, &newUrl.url)
	if err != nil {
		fmt.Printf("ERROR :%v \n", err)
		return
	}
	fmt.Printf("URL : %v \n", newUrl.id)
}
